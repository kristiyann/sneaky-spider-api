package backend

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	_ "github.com/lib/pq" // pg driver import
)

const (
	currentDatePg = "now()"
)

type Database interface {
	// users
	GetUserByAuthUserID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUsersDiscordWebhooks(ctx context.Context, paginationParams models.PaginationParams) ([]*models.UserWebhooksViewModel, error)
	InsertUser(ctx context.Context, user models.UserCreate) (*uuid.UUID, error)
	UpdateMarket(ctx context.Context, market string) error
	SetSnkrsGlobalEnabled(ctx context.Context, value bool) error
	UpdateStripeSubscription(ctx context.Context, userID *string, subscriptionID string, stripePlan *string, isCancelled bool) error
	DeleteUser(ctx context.Context) error

	// product subscriptions
	GetProductSubscriptions(ctx context.Context, paginationParams models.PaginationParams, whereStatement *string) (models.GenericPaginatedResult[models.ProductSubscriptionViewModel], error)
	GetProductSubscriptionsForAlerts(ctx context.Context, paginationParams models.PaginationParams) ([]models.ProductSubscriptionViewModel, error)
	GetActiveSubscriptionsCountByUser(ctx context.Context) (*int, error)
	InsertProductSubscription(ctx context.Context, toInsert models.ProductSubscriptionEdit) (*uuid.UUID, error)
	UpdateSubscriptionStatus(ctx context.Context, subscriptionID uuid.UUID, status string) error
	BatchUpdateSubscriptionStatus(ctx context.Context, subscriptionIDs []uuid.UUID, status string) error
	DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error
	DeleteSubscriptions(ctx context.Context) error

	// product availabilities
	GetProductAvailabilities(ctx context.Context, paginationParams models.PaginationParams, whereStatement *string) (models.GenericPaginatedResult[models.ProductAvailabilityViewModel], error)
	InsertProductAvailability(ctx context.Context, toInsert models.ProductAvailabilityEdit) (*uuid.UUID, error)
	UpdateProductAvailability(ctx context.Context, a models.ProductAvailabilityEdit) error

	// webhook urls
	InsertWebhookUrl(ctx context.Context, w models.WebhookUrlEdit) (*uuid.UUID, error)
	UpdateWebhook(ctx context.Context, discordWebhook string) error

	// logs
	InsertLog(ctx context.Context, log models.LogEdit) (*uuid.UUID, error)
}

type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type querierImpl struct {
	db  *sql.DB
	dsn string
	log *log.Logger
}

func NewQuerier(dsn string, log *log.Logger) *querierImpl {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: '%s'", err.Error())
	}

	return &querierImpl{
		log: log,
		dsn: dsn,
		db:  db,
	}
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Rows interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(dest ...interface{}) error
}

type postgresDb struct {
	querier      Querier
	db           *sql.DB
	dbConnString string
	log          *log.Logger
}

func NewPostgresDb(querier Querier, dbConn string) Database {
	return &postgresDb{
		querier:      querier,
		dbConnString: dbConn,
	}
}

func (q querierImpl) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	q.logQuery(query, args)

	return q.db.Exec(query, args...)
}

func (q querierImpl) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	q.logQuery(query, args)

	return q.db.Query(query, args...)
}

func (q querierImpl) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	q.logQuery(query, args)

	return q.db.QueryRow(query, args...)
}

func (q querierImpl) logQuery(query string, args ...interface{}) {
	if q.log != nil && util.LoadEnvVar(constants.EnvEnvironment) == constants.EnvEnvironmentValueDevelopment {
		q.log.Println(query, args)
	}
}
