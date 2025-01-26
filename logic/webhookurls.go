package logic

import (
	"context"
	"strings"

	"github.com/kristiyann/af1-spider-web-app/models"
)

const (
	discordWebhookUrlsPrefix = "https://discord.com/api/webhooks/"
)

func (l *logicImpl) UpdateWebhook(ctx context.Context, webhookUrl string) error {
	if webhookUrl != "" && !strings.HasPrefix(webhookUrl, discordWebhookUrlsPrefix) {
		return models.NewAPIError("Provided Discord Webhook URL is invalid!", 400)
	}

	err := l.db.UpdateWebhook(ctx, webhookUrl)
	if err != nil {
		return err
	}

	return nil
}
