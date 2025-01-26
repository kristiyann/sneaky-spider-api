package models

import (
	"time"

	"github.com/google/uuid"
)

// Internal public.users model
type User struct {
	PublicID           uuid.UUID
	Username           string
	Email              string
	PhoneNum           *string
	Market             string
	SnkrsGlobalEnabled bool
	StripeCustomerID   *string
	StripePlan         string
	PlanEnds           *time.Time
	DiscordWebhook     *string
}

// User model to be used for the insert process.
type UserCreate struct {
	Username         string
	Market           string
	AuthUserID       uuid.UUID
	StripeCustomerID string
}
