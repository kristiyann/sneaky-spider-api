package models

import "time"

type NikeProductViewModel struct {
	ExternalID   string                  `json:"external_id"`
	ExternalUUID string                  `json:"external_uuid"`
	Name         string                  `json:"name"`
	Link         string                  `json:"link"`
	Availability []AvailabilityViewModel `json:"availability"`
	ImageUrls    []string                `json:"image_urls"`
	Price        PriceViewModel          `json:"price"`
	Type         string                  `json:"type"`
	Colors       []ColorViewModel        `json:"colors"`
	Gender       string                  `json:"gender"`
}

type NikeProductSearchViewModel struct {
	ExternalUUID string           `json:"external_uuid"`
	ExternalID   string           `json:"external_id"`
	URL          string           `json:"url"`
	Name         string           `json:"name"`
	ImageUrls    []string         `json:"image_urls"`
	Price        PriceViewModel   `json:"price"`
	Type         string           `json:"type"`
	Colors       []ColorViewModel `json:"colors"`
	Sizes        []string         `json:"sizes"`
	Available    bool             `json:"available"`
	Market       string           `json:"market"`

	LaunchDate         time.Time `json:"-"`
	MerchProductStatus string    `json:"-"`
	AllGtinsAvailable  bool      `json:"-"`
}

type ImageUrlViewModel struct {
	URL  string `json:"url"`
	View any    `json:"view"`
}

type AvailabilityViewModel struct {
	ExternalUUID string `json:"external_uuid"`
	SizeUs       string `json:"size_us"`
	SizeLocal    string `json:"size_local"`
	Available    bool   `json:"available"`
}

type PriceViewModel struct {
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	Discounted    bool    `json:"discounted"`
	OriginalPrice float64 `json:"original_price"`
}

type ColorViewModel struct {
	StyleCode  string `json:"style_code"`
	StyleColor string `json:"style_color"`
	StyleType  string `json:"style_type"`
}
