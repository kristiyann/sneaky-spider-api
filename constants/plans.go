package constants

import "github.com/kristiyann/af1-spider-web-app/models"

const (
	PlanKeyFree         = "FREE"
	PlanKeyTest         = "prod_ObavCkCBVlq3DN"
	PlanKeyStarter      = "prod_Oari8qfnS8ySlJ"
	PlanKeyHobby        = "prod_OapIQpswmRSfnP"
	PlanKeyProfessional = "prod_OarmaNLv5uTZCA"
)

var Plans map[string]models.Plan = map[string]models.Plan{
	"FREE": {
		Name:      "Free",
		Price:     0,
		MaxAlerts: 1,
		Features:  make([]string, 0),
	},
	"prod_Oari8qfnS8ySlJ": {
		Name:      "Starter",
		Price:     499,
		MaxAlerts: 2,
		Features: []string{
			"Up to 2 concurrent alerts",
			"Email notifications",
			"Alerts never expire",
		},
	},
	"prod_OapIQpswmRSfnP": {
		Name:      "Hobby",
		Price:     799,
		MaxAlerts: 4,
		Features: []string{
			"Up to 4 concurrent alerts",
			"Email + Discord notifications",
			"Alerts never expire",
			"SNKRS monitor for Discord Webhooks",
		},
	},
	"prod_OarmaNLv5uTZCA": {
		Name:      "Professional",
		Price:     1499,
		MaxAlerts: 25,
		Features: []string{
			"Up to 25 concurrent alerts",
			"Email + Discord notifications",
			"Alerts never expire",
			"SNKRS monitor for Discord Webhooks",
		},
	},
}
