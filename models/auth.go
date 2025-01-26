package models

import (
	"time"

	"github.com/google/uuid"
)

// User model to be passed in contexts through functions.
type UserSession struct {
	ID                   uuid.UUID
	PublicID             uuid.UUID
	Username             string
	Email                string
	Picture              string
	Market               string
	PhoneNum             *string
	SnkrsGlobalEnabled   bool
	SupabaseToken        string
	SupabaseRefreshToken string
	MembershipPlan       string
	StripeCustomerID     *string
	PlanEnds             *time.Time
	DiscordWebhook       *string
}

// User session info model to be displayed to the interface.
type UserSessionInfo struct {
	ID                 uuid.UUID `json:"id"`
	PublicID           uuid.UUID `json:"public_id"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	ProfilePictureUrl  string    `json:"profile_picture_url"`
	Market             string    `json:"market"`
	MembershipPlan     string    `json:"membership_plan"`
	DiscordWebhook     *string   `json:"discord_webhook"`
	SnkrsGlobalEnabled bool      `json:"snkrs_global_enabled"`
}

type TokensPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
