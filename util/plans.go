package util

import (
	"strings"
	"time"

	"github.com/kristiyann/af1-spider-web-app/constants"
)

func GetUserMembershipPlanText(priceId string, subscriptionEndDate *time.Time) string {
	if priceId == "prod_Oari8qfnS8ySlJ" || priceId == "prod_OapIQpswmRSfnP" || priceId == "prod_OarmaNLv5uTZCA" || SubscriptionHasEnded(subscriptionEndDate) {
		return constants.Plans[priceId].Name
	}

	return "Test_Hobby"
}

func SubscriptionHasEnded(subscriptionEndDate *time.Time) bool {
	return subscriptionEndDate == nil || subscriptionEndDate.Before(time.Now())
}

func GetKeyFromPlanString(s string, isTestEnv bool) string {
	switch planString := strings.ToLower(s); planString {
	case "test_hobby":
		return "prod_ObavCkCBVlq3DN"
	case "starter":
		if isTestEnv {
			return "prod_Oari8qfnS8ySlJ"
		}
		return "prod_Oari8qfnS8ySlJ"
	case "hobby":
		if isTestEnv {
			return "prod_OapIQpswmRSfnP"
		}
		return "prod_OapIQpswmRSfnP"
	case "professional":
		if isTestEnv {
			return "prod_OarmaNLv5uTZCA"
		}
		return "prod_OarmaNLv5uTZCA"
	case "free":
		return "FREE"
	}

	return ""
}
