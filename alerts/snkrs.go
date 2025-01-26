package alerts

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

var SNKRSMarketsAddDates map[string]time.Time = map[string]time.Time{
	"BG": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"NO": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"US": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"BE": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"DE": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"GB": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"IT": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"NL": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
	"PL": time.Date(2024, 2, 21, 0, 0, 0, 0, time.Local),
}

func (a *Alerter) processSnkrs() {
	userData, err := a.getUsersDiscordWebhooksAndMarkets()
	if err != nil {
		a.sendLog("userData, err := a.getUsersDiscordWebhooksAndMarkets() "+err.Error(), nil, nil)
		panic("userData, err := a.getUsersDiscordWebhooksAndMarkets() " + err.Error())
	}

	slog.Info(fmt.Sprintf("Found webhooks count = %d", len(userData)))

	// Distinct on marketsToQueryFor
	var marketsToQueryFor = []string{
		"BG", "NO", "US", "BE", "DE", "GB", "IT", "NL", "PL",
	}
	for _, datum := range userData {
		marketsToQueryContainsCurrentMarket := false
		for _, market := range marketsToQueryFor {
			if datum.Market == market {
				marketsToQueryContainsCurrentMarket = true
			}
		}

		if !marketsToQueryContainsCurrentMarket {
			marketsToQueryFor = append(marketsToQueryFor, datum.Market)
		}
	}

	snkrsProducts, err := a.getSnkrsProducts(marketsToQueryFor)
	if err != nil {
		a.sendLog("snkrsProducts, err := a.getSnkrsProducts(marketsToQueryFor) "+err.Error(), nil, nil)
		return
	}

	params := models.PaginationParams{
		Top:  1,
		Skip: 0,
	}

	snkrsProcessorWorkers := make(chan models.NikeProductSearchViewModel)

	wg := sync.WaitGroup{}
	for i := 0; i < 25; i++ { // Batch process 25 at a time
		wg.Add(1)
		go func(params models.PaginationParams, userData []*models.UserWebhooksViewModel) {
			defer wg.Done()
			for s := range snkrsProcessorWorkers {
				err := a.processSnkrsItem(s, params, userData)
				if err != nil {
					a.sendLog("err := a.processSnkrsItem(item, params, userData): "+err.Error(), nil, nil)
					return
				}
			}
		}(params, userData)
	}

	go func() {
		for _, item := range snkrsProducts {
			snkrsProcessorWorkers <- item
		}

		close(snkrsProcessorWorkers)

		wg.Wait()
	}()

	// for _, item := range snkrsProducts {
	// 	go func(item models.NikeProductSearchViewModel, params models.PaginationParams, userData []*models.UserWebhooksViewModel) {
	// 		err := a.processSnkrsItem(item, params, userData)
	// 		if err != nil {
	// 			a.sendLog("err := a.processSnkrsItem(item, params, userData): "+err.Error(), nil, nil)
	// 			return
	// 		}
	// 	}(item, params, userData)
	// }
}

func (a *Alerter) getSnkrsProducts(marketsToQueryFor []string) ([]models.NikeProductSearchViewModel, error) {
	var mu sync.Mutex
	snkrsProducts := []models.NikeProductSearchViewModel{}
	var wg sync.WaitGroup

	for _, market := range marketsToQueryFor {
		wg.Add(1)

		go func(market string) {
			defer wg.Done()

			for skip := 0; ; skip = skip + 50 {
				params := logic.GetProductParams{
					Top:    50,
					Skip:   skip,
					Market: market,
				}
				snkrsResult, err := a.l.GetSnkrsProducts(context.Background(), params)
				if err != nil {
					a.sendLog(err.Error(), nil, nil)
				}

				if len(snkrsResult) == 0 {
					return
				}

				mu.Lock()
				snkrsProducts = append(snkrsProducts, snkrsResult...)
				mu.Unlock()
			}
		}(market)
	}

	wg.Wait()

	slog.Info(fmt.Sprintf("Found %d SNKRS items...", len(snkrsProducts)))

	return snkrsProducts, nil
}

func (a *Alerter) processSnkrsItem(item models.NikeProductSearchViewModel, params models.PaginationParams, userData []*models.UserWebhooksViewModel) error {
	if item.ExternalID == "" {
		return nil
	}

	// launchDateCutOff := time.Now().UTC().AddDate(0, -6, 0)

	whereStatement := fmt.Sprintf("WHERE a.product_external_id = '%s'", item.ExternalID)
	matchedItems, err := a.db.GetProductAvailabilities(context.Background(), params, &whereStatement)
	if err != nil {
		return err
	}

	// marketKeyExistsInAvailabilityMap := false
	// if len(*matchedItems.Data) > 0 {
	// 	_, marketKeyExistsInAvailabilityMap = (*matchedItems.Data)[0].Availability[item.Market]
	// }

	if len(*matchedItems.Data) < 1 {
		slog.Info(fmt.Sprintf("SNKRS alerts: Found no item with ID %s in the DB, attempting insert", item.ExternalID))
		// markets := []string{item.Market}
		availability := map[string]bool{
			item.Market: item.Available && strings.ToLower(item.MerchProductStatus) == "active",
		}

		var lastRecordedLaunchDate *time.Time = nil
		if (item.LaunchDate != time.Time{}) {
			lastRecordedLaunchDate = &item.LaunchDate
		}

		lastRecordedLaunchDatesMap := map[string]*time.Time{
			item.Market: lastRecordedLaunchDate,
		}

		toInsert := models.ProductAvailabilityEdit{
			ProductExternalID:      item.ExternalID,
			Vendor:                 models.VendorSNKRS,
			Available:              item.Available,
			Availability:           availability,
			LastRecordedLaunchDate: lastRecordedLaunchDatesMap,
		}

		_, err := a.db.InsertProductAvailability(context.Background(), toInsert)
		if err != nil {
			return err
		}
		if item.Available && strings.ToLower(item.MerchProductStatus) == "active" && SNKRSMarketsAddDates[item.Market].AddDate(0, 0, 7).Before(time.Now()) {
			slog.Info(fmt.Sprintf("SNKRS alerts: Item %s was available [%t] & didn't exist in the DB before", item.ExternalID, item.Available))
			err := sendDiscordWebhook(item, userData)
			if err != nil {
				return err
			}
		}
	} else if (*matchedItems.Data)[0].Availability[item.Market] != item.Available {
		now := time.Now()
		if !(*matchedItems.Data)[0].Availability[item.Market] && item.Available && strings.ToLower(item.MerchProductStatus) == "active" && (util.DatePtrEqual((*matchedItems.Data)[0].LastRecordedLaunchDate[item.Market], &now) || item.AllGtinsAvailable) {
			slog.Info(fmt.Sprintf("SNKRS alerts: Item %s was available [%t] after not being available [%t] in the DB", item.ExternalID, item.Available, (*matchedItems.Data)[0].Availability[item.Market]))
			err := sendDiscordWebhook(item, userData)
			if err != nil {
				return err
			}
		}

		// toEditMarkets := (*matchedItems.Data)[0].Market
		// if !util.StringSliceContains(toEditMarkets, item.Market) {
		// 	toEditMarkets = append(toEditMarkets, item.Market)
		// }

		availability := (*matchedItems.Data)[0].Availability
		availability[item.Market] = item.Available && strings.ToLower(item.MerchProductStatus) == "active"

		var lastRecordedLaunchDate *time.Time = nil
		if (item.LaunchDate != time.Time{}) {
			lastRecordedLaunchDate = &item.LaunchDate
		}

		lastRecordedLaunchDatesMap := (*matchedItems.Data)[0].LastRecordedLaunchDate
		lastRecordedLaunchDatesMap[item.Market] = lastRecordedLaunchDate

		toEdit := models.ProductAvailabilityEdit{
			ID:                     (*matchedItems.Data)[0].ID,
			ProductExternalID:      (*matchedItems.Data)[0].ProductExternalID,
			Vendor:                 (*matchedItems.Data)[0].Vendor,
			Available:              item.Available,
			Availability:           availability,
			LastRecordedLaunchDate: lastRecordedLaunchDatesMap,
		}
		err := a.db.UpdateProductAvailability(context.Background(), toEdit)
		if err != nil {
			return err
		}
	}

	return nil
}

func sendDiscordWebhook(item models.NikeProductSearchViewModel, userData []*models.UserWebhooksViewModel) error {
	for _, userDatum := range userData {
		if userDatum.Market == item.Market {
			slog.Info(fmt.Sprintf("SNKRS alerts: Sending webhook to %s in region %s for item %+v...", userDatum.WebhookURL, userDatum.Market, item))
			err := util.SendDiscordWebhook(userDatum.WebhookURL, item)
			if err != nil {
				return fmt.Errorf(fmt.Sprintf("SNKRS alerts: Could not send Discord webhook [%s], market=%s & webhook_url=%s", err.Error(), userDatum.Market, userDatum.WebhookURL))
			}
		}
	}

	return nil
}
