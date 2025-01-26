package logic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/stripe/stripe-go/v75"
)

// func (l *logicImpl) GetUsersDiscordWebhooks(ctx context.Context) ([]*models.UserDiscordWebhookViewModel, error) {
// 	result, err := l.db.GetUsersDiscordWebhooks(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

func (l *logicImpl) RegisterOrLoginFromProvider(ctx context.Context, authUserID uuid.UUID, username string, email string) error {
	user, err := l.db.GetUserByAuthUserID(ctx, authUserID)
	if err != nil {
		return err
	}

	if *user == (models.User{}) {
		createCustomerParams := &stripe.CustomerParams{
			Name:  &username,
			Email: &email,
		}

		c, err := l.stripeClient.Customers.New(createCustomerParams)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("Was unable to create a new Stripe customer: %s", err.Error()), http.StatusInternalServerError)
		}

		toInsert := models.UserCreate{
			Market:           "US",
			AuthUserID:       authUserID,
			Username:         username,
			StripeCustomerID: c.ID,
		}
		userID, err := l.db.InsertUser(ctx, toInsert)
		if err != nil {
			return err
		}

		webhookEntry := models.WebhookUrlEdit{
			UserID: *userID,
			Url:    "",
		}
		_, err = l.db.InsertWebhookUrl(ctx, webhookEntry)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *logicImpl) UpdateUserMarket(ctx context.Context, market string) error {
	err := l.db.UpdateMarket(ctx, market)
	if err != nil {
		return err
	}

	return nil
}

func (l *logicImpl) SetSnkrsGlobalEnabled(ctx context.Context, value bool) error {
	session := util.GetSession(ctx)
	if value == true && session.MembershipPlan != constants.PlanKeyTest && session.MembershipPlan != constants.PlanKeyProfessional {
		return models.NewAPIError("You need to be on the Professional plan in order to enable SNKRS global monitoring.", http.StatusPaymentRequired)
	}

	err := l.db.SetSnkrsGlobalEnabled(ctx, value)
	if err != nil {
		return err
	}

	return nil
}

func (l *logicImpl) UpdateUserStripeSubscription(ctx context.Context, userID *string, subscriptionID string, stripePlan *string, isCancelled bool) error {
	err := l.db.UpdateStripeSubscription(ctx, userID, subscriptionID, stripePlan, isCancelled)
	if err != nil {
		return err
	}

	return nil
}

func (l *logicImpl) DeleteUser(ctx context.Context) error {
	err := l.db.DeleteUser(ctx)
	if err != nil {
		return err
	}

	return nil
}
