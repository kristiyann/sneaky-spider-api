package models

import (
	"github.com/google/uuid"
)

type UserWebhooksViewModel struct {
	UserID     uuid.UUID
	WebhookURL string
	Market     string
}

type WebhookUrlEdit struct {
	ID     uuid.UUID
	Url    string
	UserID uuid.UUID
}
