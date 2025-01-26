package logic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/stripe/stripe-go/v75"
)

func (l *logicImpl) CreateCheckoutSession(ctx context.Context, priceId string) (*string, error) {
	currentUser := util.GetSession(ctx)

	params := &stripe.CheckoutSessionParams{
		SuccessURL:               stripe.String(util.LoadEnvVar(constants.EnvStripeReturnUrl)),
		CancelURL:                stripe.String(util.LoadEnvVar(constants.EnvStripeErrorReturnUrl)),
		Mode:                     stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		BillingAddressCollection: stripe.String(string(stripe.CheckoutSessionBillingAddressCollectionAuto)),
		//CustomerEmail:            stripe.String(currentUser.Email),
		Customer: stripe.String(*currentUser.StripeCustomerID),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price: stripe.String(priceId),
				// For metered billing, do not pass quantity
				Quantity: stripe.Int64(1),
			},
		},
		Metadata: map[string]string{
			"public_user_id": currentUser.PublicID.String(),
		},
	}

	s, err := l.stripeClient.CheckoutSessions.New(params)
	if err != nil {
		return nil, models.NewAPIError(fmt.Sprintf("Could not generate checkout link for price_id=%s: %s", priceId, err.Error()), http.StatusInternalServerError)
	}

	return &s.URL, nil
}

func (l *logicImpl) GenerateCustomerPortalLink(ctx context.Context) (*string, error) {
	currentUserSession := util.GetSession(ctx)

	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(*currentUserSession.StripeCustomerID),
		ReturnURL: stripe.String(util.LoadEnvVar(constants.EnvStripeReturnUrl)),
	}

	s, err := l.stripeBillingPortalClient.New(params)
	if err != nil {
		return nil, models.NewAPIError(fmt.Sprintf("Could not generate customer portal link for user=%s with stripe_id=%s: %s", currentUserSession.PublicID, *currentUserSession.StripeCustomerID, err.Error()), http.StatusInternalServerError)
	}

	return &s.URL, nil
}

func (l *logicImpl) GetSubscription(ctx context.Context, subscriptionID string) (*stripe.Subscription, error) {
	s, err := l.stripeClient.Subscriptions.Get(subscriptionID, nil)
	if err != nil {
		return nil, models.NewAPIError(fmt.Sprintf("could not get subscription with id=%s: %s", subscriptionID, err.Error()), http.StatusInternalServerError)
	}

	return s, nil
}
