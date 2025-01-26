package alerts

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"

	"github.com/kristiyann/af1-spider-web-app/backend"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/httpclient"
	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

type DispatchedRequest struct {
	*http.Request
	item models.ProductSubscriptionViewModel
}

type DispatchedRequestResponse struct {
	*http.Response
	err  error
	item models.ProductSubscriptionViewModel
}

type Alerter struct {
	l                         logic.Logic
	db                        backend.Database
	log                       *log.Logger
	mu                        sync.Mutex
	completedSubscriptionsMap sync.Map
	expiredSubscriptionsMap   sync.Map
	maxRequests               int
}

func NewAlerter(logic logic.Logic, db backend.Database, logger *log.Logger) Alerter {
	return Alerter{
		l:   logic,
		db:  db,
		log: logger,
	}
}

func Init() {
	dbConn := util.LoadEnvVar(constants.EnvPostgresUrl)

	dbLogger := log.New(os.Stdout, "[sql] ", log.LstdFlags|log.Llongfile)
	querier := backend.NewQuerier(dbConn, dbLogger)
	db := backend.NewPostgresDb(*querier, dbConn)

	nikeClient := httpclient.NewWithClientPool(constants.NikeBaseUrl, 3, log.New(os.Stdout, "[Nike API] ", log.LstdFlags|log.Llongfile))
	snkrsClient := httpclient.NewWithClientPool(constants.NikeBaseUrl, 3, log.New(os.Stdout, "[SNKRS API] ", log.LstdFlags|log.Llongfile))

	logic := logic.New(db, nikeClient, snkrsClient)

	alertsLogger := log.New(os.Stdout, "[alerts] ", log.LstdFlags|log.Llongfile)
	alerter := NewAlerter(logic, db, alertsLogger)

	flag.IntVar(&alerter.maxRequests, "concurrent", 300, "Maximum concurrent requests")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	go alerter.ProcessSubscriptions()
	alerter.ProcessGlobalMonitors()
}

func (a *Alerter) getUsersDiscordWebhooksAndMarkets() ([]*models.UserWebhooksViewModel, error) {
	var userData []*models.UserWebhooksViewModel
	top := 1500
	skip := 0
	for {
		retrievedData, err := a.db.GetUsersDiscordWebhooks(context.Background(), models.PaginationParams{Skip: skip, Top: top})
		if err != nil {
			a.log.Printf("[ERR] Failed retrieving discord webhook + market user data: %v", err)
			return userData, err
		}

		if len(retrievedData) > 0 {
			userData = append(userData, retrievedData...)
			skip += top
		} else {
			break
		}
	}

	return userData, nil
}
