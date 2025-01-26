package logic

import (
	"testing"
)

func TestGenerateLanguageBasedOnMarket(t *testing.T) {
	tests := []struct {
		market   string
		expected string
	}{
		{
			market:   "bg",
			expected: "en-GB",
		},
		{
			market:   "GB",
			expected: "en-GB",
		},
		{
			market:   "   it   ",
			expected: "en-GB",
		},
		{
			market:   "us",
			expected: "en",
		},
		{
			market:   "fr",
			expected: "en",
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := generateLanguageBasedOnMarket(test.market)
		if result != test.expected {
			t.Errorf("For market=%s, expected=%s, but got %s", test.market, test.expected, result)
		}
	}
}

func TestGenerateNikeProductLink(t *testing.T) {
	tests := []struct {
		market   string
		channels []string
		slug     string
		ID       string
		expected string
	}{
		{
			market:   "US",
			channels: []string{"SNKRS"},
			slug:     "product-slug",
			ID:       "product-id",
			expected: "https://nike.com/launch/t/product-slug/product-id",
		},
		{
			market:   "GB",
			channels: []string{"nike.com"},
			slug:     "another-slug",
			ID:       "another-id",
			expected: "https://nike.com/gb/t/another-slug/another-id",
		},
		{
			market:   "BG",
			channels: []string{"SNKRS", "nike.com"},
			slug:     "some-slug",
			ID:       "some-id",
			expected: "https://nike.com/bg/launch/t/some-slug/some-id",
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := generateNikeProductLink(test.market, test.channels, test.slug, test.ID)
		if result != test.expected {
			t.Errorf("For market=%s, channels=%v, slug=%s, ID=%s, expected=%s, but got %s", test.market, test.channels, test.slug, test.ID, test.expected, result)
		}
	}
}
