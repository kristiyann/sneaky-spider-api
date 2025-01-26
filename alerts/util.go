package alerts

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func (a *Alerter) sendLog(message string, stackTrace *string, logContext *string) {
	slog.Error(message)
	toInsert := models.LogEdit{
		Message:    message,
		ApiUrl:     "",
		StackTrace: stackTrace,
		Level:      models.LogLevelError,
		Source:     models.LogSourceAlerts,
		Context:    logContext,
	}
	_, err := a.db.InsertLog(context.Background(), toInsert)
	if err != nil {
		slog.Error("[alerts] Could not insert logs: " + err.Error())
	}
}

func sendProductAvailabilityEmail(toEmail string, productName string, size string, link string) error {
	// body := fmt.Sprintf("Subject: Sneaky Spider -- %s available in size %s!\r\n\r\n%s is available in size %s - %s", productName, size, productName, size, link)
	// return util.SendEmail(toEmail, "", body)
	return util.SendEmail(toEmail, fmt.Sprintf("Sneaky Spider -- %s available in size %s!", productName, size), fmt.Sprintf("%s is available in size %s - %s", productName, size, link))
}
