package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductAvailabilityViewModel struct {
	ID                     uuid.UUID             `json:"id"`
	ProductExternalID      string                `json:"product_external_id"`
	Vendor                 string                `json:"vendor"`
	CreateDate             time.Time             `json:"-"`
	CreateDateString       string                `json:"create_date"`
	Available              bool                  `json:"available"`
	AvailableSizes         []string              `json:"available_sizes"`
	Availability           map[string]bool       `json:"availability"`
	LastRecordedLaunchDate map[string]*time.Time `json:"last_recorded_launch_date"`
}

type ProductAvailabilityEdit struct {
	ID                     uuid.UUID             `json:"id"`
	ProductExternalID      string                `json:"product_external_id"`
	Vendor                 string                `json:"vendor"`
	Available              bool                  `json:"available"`
	AvailableSizes         []string              `json:"available_sizes"`
	Availability           map[string]bool       `json:"availability"`
	LastRecordedLaunchDate map[string]*time.Time `json:"last_recorded_launch_date"`
}

type ProductAvailabilityFilter struct {
	ExternalID string
}
