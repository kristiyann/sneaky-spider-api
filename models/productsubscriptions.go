package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductSubscriptionViewModel struct {
	ID                   uuid.UUID           `json:"id"`
	User                 GenericComboBoxUser `json:"user"`
	ProductExternalID    string              `json:"product_external_id"`
	Vendor               string              `json:"vendor"`
	CreateDate           time.Time           `json:"-"`
	CreateDateDisplay    string              `json:"create_date_display"`
	Size                 string              `json:"size"`
	NotificationEmail    *string             `json:"notification_email"`
	NotificationPhoneNum *string             `json:"notification_phone_num"`
	Status               string              `json:"status"`
	Market               string              `json:"market"`
	Product              any                 `json:"product"`
}

type ProductSubscriptionEdit struct {
	ID                uuid.UUID `json:"id"`
	ProductExternalID string    `json:"product_external_id"`
	Vendor            string    `json:"vendor"`
	Size              string    `json:"size"`
	Market            string    `json:"market"`
}

type ProductSubscriptionFilter struct {
	ID     uuid.UUID
	Status string
}

const (
	VendorNike  = "Nike"
	VendorSNKRS = "SNKRS"
)

func IsValidVendor(s string) bool {
	return s == VendorNike || s == VendorSNKRS
}

type ProductSubscriptionStatus string

const (
	ProductSubscriptionStatusActive    = "Active"
	ProductSubscriptionStatusCompleted = "Completed"
	ProductSubscriptionStatusCancelled = "Cancelled"
	ProductSubscriptionStatusExpired   = "Expired"
)

func IsValidProductSubscriptionStatus(s string) bool {
	return s == ProductSubscriptionStatusActive || s == ProductSubscriptionStatusCompleted || s == ProductSubscriptionStatusCancelled || s == ProductSubscriptionStatusExpired
}
