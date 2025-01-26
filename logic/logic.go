package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/backend"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/httpclient"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/billingportal/session"
	"github.com/stripe/stripe-go/v75/client"
)

type logicImpl struct {
	db                        backend.Database
	stripeClient              *client.API
	stripeBillingPortalClient session.Client
	nikeHttpClient            httpclient.HttpClient
	snkrsHttpClient           httpclient.HttpClient
}

func New(db backend.Database, nikeHttpClient httpclient.HttpClient, snkrsHttpClient httpclient.HttpClient) Logic {
	sc := &client.API{}
	sc.Init(util.LoadEnvVar(constants.EnvStripeKey), nil)

	billingPortalSessionsClient := session.Client{
		Key: util.LoadEnvVar(constants.EnvStripeKey),
		B:   stripe.GetBackend(stripe.APIBackend),
	}

	return &logicImpl{
		db:                        db,
		nikeHttpClient:            nikeHttpClient,
		snkrsHttpClient:           snkrsHttpClient,
		stripeClient:              sc,
		stripeBillingPortalClient: billingPortalSessionsClient,
	}
}

type GetProductParams struct {
	Identifier string
	Size       string
	Market     string
	Top        int
	Skip       int
}

type SearchProductParams struct {
	SearchTerm string
	Market     string
}

type Logic interface {
	// users
	RegisterOrLoginFromProvider(ctx context.Context, authUserID uuid.UUID, username string, email string) error
	UpdateUserMarket(ctx context.Context, market string) error
	SetSnkrsGlobalEnabled(ctx context.Context, value bool) error
	UpdateUserStripeSubscription(ctx context.Context, userID *string, subscriptionID string, stripePlan *string, isCancelled bool) error
	DeleteUser(ctx context.Context) error
	//GetUsersDiscordWebhooks(ctx context.Context) ([]*models.UserDiscordWebhookViewModel, error)

	// nike
	GetNikeProduct(ctx context.Context, params GetProductParams) (*models.NikeProductViewModel, error)
	SearchNikeProducts(ctx context.Context, params SearchProductParams) ([]models.NikeProductSearchViewModel, error)

	// snkrs
	GetSnkrsProducts(ctx context.Context, params GetProductParams) ([]models.NikeProductSearchViewModel, error)

	// subscriptions
	GetSubscriptions(ctx context.Context, paginationParams models.PaginationParams, filter models.ProductSubscriptionFilter, expand string) (models.GenericPaginatedResult[models.ProductSubscriptionViewModel], error)
	GetSubscriptionsForAlerts(ctx context.Context, paginationParams models.PaginationParams) ([]models.ProductSubscriptionViewModel, error)
	InsertSubscription(ctx context.Context, toInsert models.ProductSubscriptionEdit) (*uuid.UUID, error)
	UpdateSubscriptionStatus(ctx context.Context, ID uuid.UUID, status string) error
	BatchUpdateSubscriptionStatus(ctx context.Context, IDs []uuid.UUID, status string) error
	DeleteSubscriptions(ctx context.Context, ID uuid.UUID) error

	// webhook_urls
	UpdateWebhook(ctx context.Context, discordWebhook string) error

	// stripe
	CreateCheckoutSession(ctx context.Context, priceId string) (*string, error)
	GenerateCustomerPortalLink(ctx context.Context) (*string, error)
	GetSubscription(ctx context.Context, subscriptionID string) (*stripe.Subscription, error)
}
