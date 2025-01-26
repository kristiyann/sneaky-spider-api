package alerts

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func (a *Alerter) subscriptionsDispatcher(reqChan chan DispatchedRequest, subscriptions []models.ProductSubscriptionViewModel) {
	defer close(reqChan)

	date2WeeksAgo := time.Now().AddDate(0, 0, -14)
	for i := 0; i < len(subscriptions); i++ {
		// If the subscription is 2 weeks old & the user is on the free plan it becomes expired
		if date2WeeksAgo.After(subscriptions[i].CreateDate) && strings.ToLower(*subscriptions[i].User.MembershipPlan) == "free" {
			a.expiredSubscriptionsMap.Store(subscriptions[i].ID, true)
		} else {
			getParams := logic.GetProductParams{
				Identifier: subscriptions[i].ProductExternalID,
				Size:       subscriptions[i].Size,
				Market:     subscriptions[i].Market,
			}
			url := fmt.Sprintf("%s%s", constants.NikeBaseUrl, logic.GetNikeProductRequestUrl(getParams))
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				a.sendLog("req, err := http.NewRequest(http.MethodGet, url, nil): "+err.Error(), nil, util.StringPtr("subscriptionsDispatcher"))
			}

			reqChan <- DispatchedRequest{
				Request: req,
				item:    subscriptions[i],
			}
		}
	}
}

func (a *Alerter) workerPool(reqChan chan DispatchedRequest, respChan chan DispatchedRequestResponse) {
	t := &http.Transport{}
	for i := 0; i < a.maxRequests; i++ {
		go worker(t, reqChan, respChan)
	}
}

func worker(t *http.Transport, reqChan chan DispatchedRequest, respChan chan DispatchedRequestResponse) {
	for req := range reqChan {
		resp, err := t.RoundTrip(req.Request)
		r := DispatchedRequestResponse{resp, err, req.item}
		respChan <- r
	}
}

func (a *Alerter) subscriptionsConsumer(respChan chan DispatchedRequestResponse, subscriptionsLen int) (int64, int64, error) {
	var (
		conns int64
		size  int64
	)
	for conns < int64(subscriptionsLen) {
		select {
		case r, ok := <-respChan:
			if ok {
				completed := false
				if r.err != nil {
					return 0, 0, r.err
				} else {
					size += r.ContentLength

					switch vendor := r.item.Vendor; vendor {
					case models.VendorNike:
						{
							completed = a.processNike(r)
						}
						break
					case models.VendorSNKRS:
						break
					default:
						break
					}

					a.completedSubscriptionsMap.Store(r.item.ID, completed)
				}

				conns++
			}
		}
	}

	return conns, size, nil
}
