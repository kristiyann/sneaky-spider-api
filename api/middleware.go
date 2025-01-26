package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/backend"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/nedpals/supabase-go"
)

const authorizationHeader = "Authorization"

func authorizationMiddleware(next func(rw http.ResponseWriter, r *http.Request), supabaseClient *supabase.Client, db backend.Database) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get(authorizationHeader), "Bearer ")
		if len(authHeader) != 2 {
			slog.Debug("Malformed token")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		accessToken := authHeader[1]

		secretBytes := []byte(util.LoadEnvVar("SUPABASE_SECRET"))

		// Parse the JWT token
		_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret key for validation
			return secretBytes, nil
		})
		if err != nil {
			slog.Error("Issue validating token: " + err.Error())
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		currentUser, err := supabaseClient.Auth.User(context.Background(), accessToken)
		if err != nil {
			handleApiError(rw, models.NewAPIError("Problem retrieving user from Supabase: "+err.Error(), http.StatusUnauthorized))
			return
		}

		currentUserID, err := uuid.Parse(currentUser.ID)
		if err != nil {
			handleApiError(rw, models.NewAPIError(fmt.Sprintf("Problem decoding userID=%s", currentUser.ID), http.StatusInternalServerError))
			return
		}

		publicUserResult, err := db.GetUserByAuthUserID(context.Background(), currentUserID)
		if err != nil {
			handleApiError(rw, models.NewAPIError("Authorization Middleware: Retrieving public.users entry from supabase: "+err.Error(), http.StatusFailedDependency))
			return
		}

		session := models.UserSession{
			ID:                 currentUserID,
			PublicID:           publicUserResult.PublicID,
			Username:           publicUserResult.Username,
			Email:              currentUser.Email,
			Picture:            currentUser.UserMetadata[supabaseRawUserDataKey_Picture].(string),
			Market:             publicUserResult.Market,
			PhoneNum:           publicUserResult.PhoneNum,
			SnkrsGlobalEnabled: publicUserResult.SnkrsGlobalEnabled,
			SupabaseToken:      accessToken,
			MembershipPlan:     publicUserResult.StripePlan,
			StripeCustomerID:   publicUserResult.StripeCustomerID,
			PlanEnds:           publicUserResult.PlanEnds,
			DiscordWebhook:     publicUserResult.DiscordWebhook,
		}

		slog.Info("Proceeding with request from user [" + session.ID.String() + "]")

		r = r.WithContext(context.WithValue(r.Context(), constants.ContextKeyUserSession, session))

		next(rw, r)
	})
}

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if !contentTypeIsJson(rw, r) {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusUnsupportedMediaType)
			err := json.NewEncoder(rw).Encode("Request expected application/json")
			if err != nil {
				slog.Error("writing json")
				return
			}
			return
		}

		next.ServeHTTP(rw, r)
	})
}

func requestLoggingMiddleware(next func(rw http.ResponseWriter, r *http.Request), db backend.Database) func(rw http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// session := util.GetSession(r.Context())

		requestId := uuid.New().String()
		slog.Info("Begin API Request [" + requestId + "]: " + r.RequestURI)

		reqHeadersBytes, err := json.Marshal(r.Header)
		if err != nil {
			slog.Error("Could not unmarshal header")
		}

		slog.Info("[" + requestId + "] Headers: " + string(reqHeadersBytes))

		defer func() {
			if err := recover(); err != nil {
				log.Println("Stack Trace: ", string(debug.Stack()))
				handleApiError(rw, models.NewAPIError(fmt.Sprintf("%+v", err), http.StatusInternalServerError))
			}
		}()

		next(rw, r)
	})
}
