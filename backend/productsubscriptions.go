package backend

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/lib/pq"
)

const (
	getSubscriptionsStatement = `SELECT s.id, s.vendor, s.product_external_id, s.create_date, s.size, s.notification_email, u.phone_num as phone_num, s.status, s.market, u.id as user_id, u.username as username, u.stripe_plan as membership_plan, wu.url FROM product_subscriptions s
								 INNER JOIN users AS u 
								 ON s.user_id = u.id
								 INNER JOIN webhook_urls AS wu
								 ON wu.user_id = u.id`
	insertSubscriptionStatement            = `INSERT INTO product_subscriptions (user_id, vendor, product_external_id, create_date, size, status, market, notification_email) VALUES ($1, $2, $3, ` + currentDatePg + `, $4, 'Active', $5, $6) RETURNING id;`
	updateSubscriptionStatusStatement      = `UPDATE product_subscriptions SET status = $3, status_last_change_date = ` + currentDatePg + ` WHERE id = $1 AND user_id = $2;`
	updateSubscriptionStatusStatementBatch = `UPDATE product_subscriptions SET status = $2, status_last_change_date = ` + currentDatePg + ` WHERE id=ANY($1);`
	deleteSubscriptionStatement            = `DELETE FROM product_subscriptions WHERE id = $1 AND user_id = $2;`
	deleteSubscriptionsStatement           = `DELETE FROM product_subscriptions WHERE user_id = $1;`
)

func (db *postgresDb) GetProductSubscriptions(ctx context.Context, paginationParams models.PaginationParams, whereStatement *string) (models.GenericPaginatedResult[models.ProductSubscriptionViewModel], error) {
	currentUser := util.GetSession(ctx)

	var list []models.ProductSubscriptionViewModel = []models.ProductSubscriptionViewModel{}
	var result models.GenericPaginatedResult[models.ProductSubscriptionViewModel]
	var filter string

	result.Data = &list
	result.Top = paginationParams.Top
	result.Skip = paginationParams.Skip

	if whereStatement == nil {
		filter = "WHERE s.user_id = $1"
	} else {
		filter = "WHERE s.user_id = $1 " + *whereStatement
	}

	sqlStatement := generatePaginatedSqlStatement(getSubscriptionsStatement, paginationParams.Top, paginationParams.Skip, "create_date DESC", filter)

	rows, err := db.querier.Query(ctx, sqlStatement, currentUser.PublicID.String())
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbObj models.ProductSubscriptionViewModel
		var user models.GenericComboBoxUser
		var createDateString string
		err := rows.Scan(
			&dbObj.ID,
			&dbObj.Vendor,
			&dbObj.ProductExternalID,
			&createDateString,
			&dbObj.Size,
			&dbObj.NotificationEmail,
			&dbObj.NotificationPhoneNum,
			&dbObj.Status,
			&dbObj.Market,
			&user.Value,
			&user.Text,
			&user.MembershipPlan,
			&user.DiscordWebhookUrl,
			&result.Count,
		)
		if err != nil {
			return result, models.NewAPIError("failed scanning row: "+err.Error(), http.StatusInternalServerError)
		}

		parsedCreateDate, err := time.Parse(time.RFC3339, createDateString)
		if err != nil {
			return result, models.NewAPIError("time.Parse(time.RFC3339, createDateString): "+err.Error(), http.StatusInternalServerError)
		}
		dbObj.CreateDate = parsedCreateDate
		dbObj.CreateDateDisplay = createDateString
		dbObj.User = user
		list = append(list, dbObj)
	}

	return result, nil
}

func (db *postgresDb) GetProductSubscriptionsForAlerts(ctx context.Context, paginationParams models.PaginationParams) ([]models.ProductSubscriptionViewModel, error) {
	var list []models.ProductSubscriptionViewModel = []models.ProductSubscriptionViewModel{}

	sqlStatement := `SELECT s.id, s.vendor, s.product_external_id, s.create_date, s.size, s.notification_email, u.phone_num as phone_num, s.status, s.market, u.id as user_id, u.username as username, u.stripe_plan as membership_plan, wu.url FROM product_subscriptions s
						INNER JOIN users AS u 
						ON s.user_id = u.id
						INNER JOIN webhook_urls AS wu
						ON wu.user_id = u.id 
						WHERE s.status = 'Active' 
						ORDER BY s.create_date OFFSET $1 LIMIT $2;`

	rows, err := db.querier.Query(ctx, sqlStatement, paginationParams.Skip, paginationParams.Top)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbObj models.ProductSubscriptionViewModel
		var user models.GenericComboBoxUser
		var createDateString string
		err := rows.Scan(
			&dbObj.ID,
			&dbObj.Vendor,
			&dbObj.ProductExternalID,
			&createDateString,
			&dbObj.Size,
			&dbObj.NotificationEmail,
			&dbObj.NotificationPhoneNum,
			&dbObj.Status,
			&dbObj.Market,
			&user.Value,
			&user.Text,
			&user.MembershipPlan,
			&user.DiscordWebhookUrl,
		)
		if err != nil {
			return list, models.NewAPIError("GetProductSubscriptionsForAlerts: Could not scan SQL row: "+err.Error(), http.StatusInternalServerError)
		}

		parsedCreateDate, err := time.Parse(time.RFC3339, createDateString)
		if err != nil {
			return list, models.NewAPIError("GetProductSubscriptionsForAlerts: parsedCreateDate, err := time.Parse(time.RFC3339, createDateString): "+err.Error(), http.StatusInternalServerError)
		}
		dbObj.CreateDate = parsedCreateDate
		dbObj.CreateDateDisplay = createDateString
		dbObj.User = user
		list = append(list, dbObj)
	}

	return list, nil
}

func (db *postgresDb) GetActiveSubscriptionsCountByUser(ctx context.Context) (*int, error) {
	session := util.GetSession(ctx)
	var count int

	sqlStatement := `SELECT COUNT(*) FROM product_subscriptions WHERE user_id = $1 AND status = 'Active';`

	rows, err := db.querier.Query(ctx, sqlStatement, session.PublicID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&count,
		)
		if err != nil {
			return nil, models.NewAPIError("GetActiveSubscriptionsCountByUser: Could not scan SQL row: "+err.Error(), http.StatusInternalServerError)
		}
	}

	return &count, nil
}

func (db *postgresDb) InsertProductSubscription(ctx context.Context, toInsert models.ProductSubscriptionEdit) (*uuid.UUID, error) {
	session := util.GetSession(ctx)
	var id uuid.UUID

	sqlStatement := insertSubscriptionStatement

	var market string = session.Market
	if toInsert.Market != "" {
		market = toInsert.Market
	}

	err := db.querier.QueryRow(
		ctx,
		sqlStatement,
		session.PublicID.String(),
		toInsert.Vendor,
		toInsert.ProductExternalID,
		toInsert.Size,
		market,
		session.Email,
	).Scan(&id)
	if err != nil {
		switch err {
		default:
			return nil, err
		}
	}

	return &id, nil
}

func (db *postgresDb) UpdateSubscriptionStatus(ctx context.Context, subscriptionID uuid.UUID, status string) error {
	session := util.GetSession(ctx)
	result, err := db.querier.Exec(ctx, updateSubscriptionStatusStatement, subscriptionID.String(), session.PublicID, status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return models.NewAPIError("Failed to update subscription status", http.StatusBadRequest)
	}

	return nil
}

func (db *postgresDb) BatchUpdateSubscriptionStatus(ctx context.Context, subscriptionIDs []uuid.UUID, status string) error {
	result, err := db.querier.Exec(ctx, updateSubscriptionStatusStatementBatch, pq.Array(subscriptionIDs), status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return models.NewAPIError("No rows were affected", http.StatusBadRequest)
	}

	return nil
}

func (db *postgresDb) DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error {
	session := util.GetSession(ctx)
	result, err := db.querier.Exec(ctx, deleteSubscriptionStatement, subscriptionID.String(), session.PublicID.String())
	if err != nil {
		return models.NewAPIError("Failed to delete subscriptions: "+err.Error(), http.StatusInternalServerError)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return models.NewAPIError("Failed to delete subscription - does not exist", http.StatusBadRequest)
	}

	return nil
}

func (db *postgresDb) DeleteSubscriptions(ctx context.Context) error {
	session := util.GetSession(ctx)
	_, err := db.querier.Exec(ctx, deleteSubscriptionsStatement, session.PublicID.String())
	if err != nil {
		return models.NewAPIError("Failed to delete subscriptions: "+err.Error(), http.StatusInternalServerError)
	}

	return nil
}
