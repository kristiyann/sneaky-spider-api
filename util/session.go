package util

import (
	"context"

	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
)

func GetSession(ctx context.Context) models.UserSession {
	session, ok := ctx.Value(constants.ContextKeyUserSession).(models.UserSession)
	if !ok {
		return models.UserSession{}
	}

	return session
}
