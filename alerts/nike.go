package alerts

import (
	"log"
	"log/slog"

	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
)

func (a *Alerter) processNike(r DispatchedRequestResponse) bool {
	completed := false

	// convert to viewmodel
	responseModel := logic.DecodeNikeProductResponse2(r.Response, false)
	r.Body.Close()
	if responseModel == nil || responseModel == &(models.NikeProductResponse2{}) {
		log.Println("moved on to option 2")
		responseModel := logic.DecodeNikeProductResponse1(r.Response)
		log.Println(responseModel)

		if len(responseModel.Objects) > 0 {
			result := logic.ConvertNikeProductResposeToNikeProductViewModel(responseModel)
			for _, availability := range result.Availability {
				if availability.SizeLocal == r.item.Size && availability.Available {
					completed = a.dispatchNikeAlerts(r, result)
					// // 2. set subscription status as completed
					// completed = true
				}
			}
		}
	} else if len(responseModel.Objects) > 0 {
		result := logic.ConvertNikeProductRespose2ToNikeProductViewModel(responseModel)
		for _, availability := range result.Availability {
			if availability.SizeLocal == r.item.Size && availability.Available {
				completed = a.dispatchNikeAlerts(r, result)
				// // 2. set subscription status as completed
				// completed = true
			}
		}
	}

	return completed
}

func (a *Alerter) dispatchNikeAlerts(r DispatchedRequestResponse, result *models.NikeProductViewModel) bool {
	// 1. alert user (item.email, item.phone)
	slog.Info("Sending out Email for " + result.Name)
	if r.item.NotificationEmail != nil {
		err := sendProductAvailabilityEmail(*r.item.NotificationEmail, result.Name, r.item.Size, result.Link)
		if err != nil {
			a.sendLog(err.Error(), nil, nil)
			return false
		}
	}

	// if r.item.NotificationPhoneNum != nil {
	// 	// send sms msg
	// }

	// if r.item.User.DiscordWebhookUrl != nil && *(r.item.User.DiscordWebhookUrl) != "" {
	// 	// err := util.SendDiscordWebhook(*r.item.User.DiscordWebhookUrl, r.item)
	// 	// if err != nil {
	// 	// 	a.sendLog(err.Error(), nil, nil)
	// 	// 	return completed
	// 	// }
	// }

	return true
}
