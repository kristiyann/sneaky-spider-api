package util

import (
	"github.com/kristiyann/af1-spider-web-app/constants"
)

func GetCorsAllowedOrigins() []string {
	environment := LoadEnvVar("ENVIRONMENT")

	if environment == "development" {
		return constants.CorsAllowedOriginsDevelopment
	} else if environment == "production" {
		return constants.CorsAllowedOriginsProduction
	}

	return []string{}
}
