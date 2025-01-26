package backend

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/lib/pq"
)

func (db *postgresDb) GetUserByAuthUserID(ctx context.Context, authUserID uuid.UUID) (*models.User, error) {
	var user models.User

	sqlStatement := `SELECT u.id, u.market, u.phone_num, u.username, u.stripe_customer_id, u.stripe_plan, u.plan_ends, u.snkrs_global_enabled, w.url FROM users AS u 
						INNER JOIN webhook_urls w ON w.user_id = u.id 
						WHERE auth_user_id = $1;`
	rows, err := db.querier.Query(ctx, sqlStatement, authUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var planEndsPg *string
		err := rows.Scan(
			&user.PublicID,
			&user.Market,
			&user.PhoneNum,
			&user.Username,
			&user.StripeCustomerID,
			&user.StripePlan,
			&planEndsPg,
			&user.SnkrsGlobalEnabled,
			&user.DiscordWebhook,
		)
		if err != nil {
			return nil, models.NewAPIError("Failed to retrieve user by auth_user_id: "+err.Error(), http.StatusInternalServerError)
		}

		if planEndsPg != nil {
			parsedTime, err := time.Parse(time.RFC3339, *planEndsPg)
			if err != nil {
				return nil, models.NewAPIError("time.Parse(time.RFC3339, *planEndsPg): "+err.Error(), http.StatusInternalServerError)
			}

			user.PlanEnds = &parsedTime
		}
	}

	return &user, nil
}

func (db *postgresDb) GetUsersDiscordWebhooks(ctx context.Context, paginationParams models.PaginationParams) ([]*models.UserWebhooksViewModel, error) {
	var result = []*models.UserWebhooksViewModel{}

	sqlStatement := fmt.Sprintf(
		`SELECT u.id, w.url, u.market FROM users as u
			INNER JOIN webhook_urls w
			ON w.user_id = u.id 
			WHERE w.url IS NOT NULL 
			AND snkrs_global_enabled = true
			AND plan_ends > now()
			AND (stripe_plan = '%s' OR stripe_plan = '%s')
			ORDER BY u.id OFFSET $1 LIMIT $2;`, constants.PlanKeyProfessional, constants.PlanKeyTest)
	rows, err := db.querier.Query(ctx, sqlStatement, paginationParams.Skip, paginationParams.Top)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbObj models.UserWebhooksViewModel
		err := rows.Scan(
			&dbObj.UserID,
			&dbObj.WebhookURL,
			&dbObj.Market,
		)
		if err != nil {
			return nil, models.NewAPIError("Failed to retrieve Webhook URLs: "+err.Error(), http.StatusInternalServerError)
		}

		result = append(result, &dbObj)
	}

	return result, nil
}

func (db *postgresDb) InsertUser(ctx context.Context, user models.UserCreate) (*uuid.UUID, error) {
	var id uuid.UUID

	sqlStatement := "INSERT INTO users (market, phone_num, auth_user_id, username, stripe_customer_id) VALUES ($1, '', $2, $3, $4) RETURNING id"

	err := db.querier.QueryRow(ctx, sqlStatement,
		user.Market,
		user.AuthUserID,
		user.Username,
		user.StripeCustomerID,
	).Scan(&id)
	if err != nil {
		return nil, models.NewAPIError("Failed to add user: "+err.Error(), http.StatusInternalServerError)
	}

	return &id, nil
}

func (db *postgresDb) UpdateMarket(ctx context.Context, market string) error {
	session := util.GetSession(ctx)

	sqlStatement := "UPDATE users SET market = $3 WHERE id = $1 AND auth_user_id = $2;"
	result, err := db.querier.Exec(ctx, sqlStatement, session.PublicID, session.ID, market)
	if err != nil {
		return models.NewAPIError("Failed to update user region: "+err.Error(), http.StatusInternalServerError)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.NewAPIError("Failed to update user region", http.StatusBadRequest)
	}

	return nil
}

func (db *postgresDb) SetSnkrsGlobalEnabled(ctx context.Context, value bool) error {
	session := util.GetSession(ctx)

	sqlStatement := "UPDATE users SET snkrs_global_enabled = $3 WHERE id = $1 AND auth_user_id = $2;"
	result, err := db.querier.Exec(ctx, sqlStatement, session.PublicID, session.ID, value)
	if err != nil {
		return models.NewAPIError("Failed to update snkrs_global_enabled: "+err.Error(), http.StatusInternalServerError)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.NewAPIError("No rows were affected while updating snkrs_global_enabled.", http.StatusBadRequest)
	}

	return nil
}

func (db *postgresDb) UpdateStripeSubscription(ctx context.Context, userID *string, subscriptionID string, stripePlan *string, isCancelled bool) error {
	sqlStatement := "UPDATE users SET stripe_subscription_id = $2, plan_ends = $4, stripe_plan = $3 WHERE id = $1;"
	var result sql.Result

	todayPlusMonth := (time.Now().AddDate(0, 1, 0))
	var planEnds interface{} = &todayPlusMonth
	subscriptionNewValue := subscriptionID
	if isCancelled {
		subscriptionNewValue = ""
		planEnds = pq.NullTime{}
	}

	if userID == nil {
		sqlStatement = "UPDATE users SET stripe_subscription_id = $2, plan_ends = $4, stripe_plan = $3 WHERE stripe_subscription_id = $1;"

		res, err := db.querier.Exec(ctx, sqlStatement, subscriptionID, subscriptionNewValue, *stripePlan, planEnds)
		if err != nil {
			return models.NewAPIError("Failed to update user stripe subscription: "+err.Error(), http.StatusInternalServerError)
		}

		result = res
	} else {
		res, err := db.querier.Exec(ctx, sqlStatement, userID, subscriptionID, *stripePlan, planEnds)
		if err != nil {
			return models.NewAPIError("Failed to update user stripe subscription: "+err.Error(), http.StatusInternalServerError)
		}

		result = res
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.NewAPIError("rowsAffected, err := result.RowsAffected()", http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		return models.NewAPIError("Failed to update user stripe subscription: 0 rows affected", http.StatusBadRequest)
	}

	return nil
}

func (db *postgresDb) DeleteUser(ctx context.Context) error {
	session := util.GetSession(ctx)

	err := db.DeleteSubscriptions(ctx)
	if err != nil {
		return models.NewAPIError(fmt.Sprintf("Failed to delete user's product subscriptions: %s", err.Error()), http.StatusInternalServerError)
	}

	err = db.DeleteLogs(ctx)
	if err != nil {
		return models.NewAPIError(fmt.Sprintf("Failed to delete user's logs: %s", err.Error()), http.StatusInternalServerError)
	}

	_, err = db.querier.Exec(ctx, `DELETE FROM users WHERE id = $1;`, session.PublicID.String())
	if err != nil {
		return models.NewAPIError("Failed to delete user info: "+err.Error(), http.StatusInternalServerError)
	}

	return nil
}
