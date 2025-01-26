package alerts

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func (a *Alerter) ProcessSubscriptions() {
	var runCount int64 = 1

	for {
		var subscriptions []models.ProductSubscriptionViewModel
		// Retrieve active subscriptions with pagination
		top := 1500
		skip := 0
		for {
			retrievedSubscriptions, err := a.db.GetProductSubscriptionsForAlerts(context.Background(), models.PaginationParams{Skip: skip, Top: top})
			if err != nil {
				a.sendLog(fmt.Sprintf("Could not retrieve subscriptions: %v", err), nil, util.StringPtr("ProcessSubscriptions"))
				// a.log.Println(fmt.Sprintf("[ERR] failed retrieving subscriptions: %v", err))
				return
			}

			if len(retrievedSubscriptions) > 0 {
				subscriptions = append(subscriptions, retrievedSubscriptions...)
				skip += top
			} else {
				break
			}
		}

		a.log.Printf("Run %d for %d items \n", runCount, len(subscriptions))

		// var wg sync.WaitGroup

		// for _, item := range subscriptions {
		// 	wg.Add(1)
		// 	// dispatcher
		// 	go func(subscription models.ProductSubscriptionViewModel) {
		// 		defer wg.Done()

		// 		// If the subscription is 2 weeks old & the user is on the free plan it becomes expired
		// 		if (time.Now().AddDate(0, 0, -14)).After(subscription.CreateDate) && strings.ToLower(*subscription.User.MembershipPlan) == "free" {
		// 			a.expiredSubscriptionsMap.Store(subscription.ID, true)
		// 		} else {
		// 			a.analyzeNikeProduct(subscription)
		// 		}
		// 	}(item)
		// }

		// wg.Wait()

		reqChan := make(chan DispatchedRequest)
		respChan := make(chan DispatchedRequestResponse)
		start := time.Now()

		go a.subscriptionsDispatcher(reqChan, subscriptions)
		go a.workerPool(reqChan, respChan)
		conns, size, err := a.subscriptionsConsumer(respChan, len(subscriptions))
		if err != nil {
			a.sendLog(fmt.Sprintf("conns, size, err := a.subscriptionsConsumer(respChan, len(subscriptions)): %s", err.Error()), nil, util.StringPtr("ProcessSubscriptions"))
			return
		}

		if conns > 0 {
			took := time.Since(start)
			ns := took.Nanoseconds()
			av := ns / conns
			average, err := time.ParseDuration(fmt.Sprintf("%d", av) + "ns")
			if err != nil {
				a.sendLog(err.Error(), nil, nil)
			}
			fmt.Printf("Connections:\t%d\nConcurrent:\t%d\nTotal size:\t%d bytes\nTotal time:\t%s\nAverage time:\t%s\n", conns, a.maxRequests, size, took, average)
		}

		var expiredSubscriptionsIDs []uuid.UUID
		a.expiredSubscriptionsMap.Range(func(id, value interface{}) bool {
			expiredSubscriptionsIDs = append(expiredSubscriptionsIDs, id.(uuid.UUID))
			return true
		})

		var completedSubscriptionsIDs []uuid.UUID
		a.completedSubscriptionsMap.Range(func(id, value interface{}) bool {
			if completed, ok := value.(bool); ok && completed {
				completedSubscriptionsIDs = append(completedSubscriptionsIDs, id.(uuid.UUID))
			}
			return true
		})

		// Update expired subs
		a.mu.Lock()
		if len(expiredSubscriptionsIDs) > 0 {
			err := a.db.BatchUpdateSubscriptionStatus(context.Background(), expiredSubscriptionsIDs, models.ProductSubscriptionStatusExpired)
			if err != nil {
				a.sendLog(fmt.Sprintf("Failed updating subscriptions to expired status, count:%d: %s \n", len(expiredSubscriptionsIDs), err.Error()), nil, util.StringPtr("ProcessSubscriptions"))
				return
			}
			a.expiredSubscriptionsMap = sync.Map{}
		}
		a.mu.Unlock()

		// Update completed subs
		a.mu.Lock()
		if len(completedSubscriptionsIDs) > 0 {
			err := a.db.BatchUpdateSubscriptionStatus(context.Background(), completedSubscriptionsIDs, models.ProductSubscriptionStatusCompleted)
			if err != nil {
				a.sendLog(fmt.Sprintf("Failed updating subscriptions to completed status, count:%d: %s \n", len(completedSubscriptionsIDs), err.Error()), nil, util.StringPtr("ProcessSubscriptions"))
				return
			}
			a.completedSubscriptionsMap = sync.Map{}
		}
		a.mu.Unlock()

		runCount++
	}
}

func (a *Alerter) ProcessGlobalMonitors() {
	for {
		var globalRunCount int64 = 0

		a.processSnkrs()

		globalRunCount++
	}
}
