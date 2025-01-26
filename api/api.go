package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/backend"
	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/nedpals/supabase-go"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const defaultTop = 10

type Server struct {
	listenAddr     string
	logic          logic.Logic
	supabaseClient *supabase.Client
	db             backend.Database
	log            *log.Logger
}

func NewServer(listenAddr string, logic logic.Logic, supabaseClient *supabase.Client, db backend.Database, logger *log.Logger) *Server {
	return &Server{
		listenAddr:     listenAddr,
		logic:          logic,
		db:             db,
		supabaseClient: supabaseClient,
		log:            logger,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()

	// health check
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		writeJson(rw, http.StatusOK, "Sneaky Spider is running!")
	}).Methods(http.MethodGet)
	router.HandleFunc("/health-check", func(rw http.ResponseWriter, r *http.Request) {
		writeJson(rw, http.StatusOK, "Sneaky Spider is running!")
	}).Methods(http.MethodGet)

	// users
	router.HandleFunc("/api/users", s.authorize(s.GetCurrentUser)).Methods(http.MethodGet)
	router.HandleFunc("/api/users/google/handle-signin", s.createHandler(s.HandleGoogleSignIn)).Methods(http.MethodPost)
	router.HandleFunc("/api/users/signout", s.authorize(s.SignOut)).Methods(http.MethodPost)
	router.HandleFunc("/api/users/refresh-token", s.createHandler(s.RefreshToken)).Methods(http.MethodPost)
	router.HandleFunc("/api/users/update-market", s.authorize(s.UpdateUserMarket)).Methods(http.MethodPut)
	router.HandleFunc("/api/users/monitors/snkrs", s.authorize(s.SetSnkrsGlobalEnabled)).Methods(http.MethodPut)
	router.HandleFunc("/api/users", s.authorize(s.DeleteUserInfo)).Methods(http.MethodDelete)

	// vendors
	router.HandleFunc("/api/vendors", s.createHandler(s.GetVendors)).Methods(http.MethodGet)

	// nike
	router.HandleFunc("/api/nike", s.createHandler(s.GetNikeProduct)).Methods(http.MethodGet)
	router.HandleFunc("/api/nike/search", s.createHandler(s.SearchNikeProducts)).Methods(http.MethodGet)

	// snkrs
	router.HandleFunc("/api/snkrs/search", s.createHandler(s.SearchSnkrsProducts)).Methods(http.MethodGet)

	// subscriptions
	router.HandleFunc("/api/subscriptions", s.authorize(s.GetSubscriptions)).Methods(http.MethodGet)
	router.HandleFunc("/api/subscriptions", s.authorize(s.InsertSubscription)).Methods(http.MethodPost)
	router.HandleFunc("/api/subscriptions/update-status", s.authorize(s.UpdateSubscriptionStatus)).Methods(http.MethodPut)
	router.HandleFunc("/api/subscriptions", s.authorize(s.DeleteSubscription)).Methods(http.MethodDelete)

	// webhook_urls
	router.HandleFunc("/api/webhook-urls", s.authorize(s.UpdateWebhookUrl)).Methods(http.MethodPut)

	// payments
	router.HandleFunc("/api/payments/checkout-link", s.authorize(s.GetCheckoutLink)).Methods(http.MethodGet)
	router.HandleFunc("/api/payments/customer-portal-link", s.authorize(s.GetCustomerPortalLink)).Methods(http.MethodGet)
	router.HandleFunc("/api/payments/stripe/webhook", s.createHandler(s.StripeWebhook)).Methods(http.MethodPost)

	// // universal middleware
	// router.Use(requestLoggingMiddleware)

	// CORS
	headersOk := gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Origin", "Authorization"})
	originsOk := gohandlers.AllowedOrigins(util.GetCorsAllowedOrigins())
	methodsOk := gohandlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})
	credentialsOk := gohandlers.AllowCredentials()

	slog.Info("Server running on address: " + s.listenAddr)
	http.ListenAndServe(s.listenAddr, gohandlers.CORS(headersOk, originsOk, methodsOk, credentialsOk)(router))
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func (s *Server) createHandler(f handlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		session := util.GetSession(r.Context())
		var userID *uuid.UUID = nil
		if session.ID != uuid.Nil {
			userID = &session.PublicID
		}

		defer func() {
			if err := recover(); err != nil {
				log.Println("Stack Trace: ", string(debug.Stack()))
				s.logError(r, fmt.Sprintf("%+v", err), http.StatusInternalServerError, util.StringPtr(string(debug.Stack())), userID)
				handleApiError(rw, models.NewAPIError(fmt.Sprintf("%+v", err), http.StatusInternalServerError))
			}
		}()

		err := f(rw, r)
		if err != nil {
			var status int
			switch v := err.(type) {
			case *models.APIError:
				status = v.HttpStatus
			default:
				status = 0
			}

			s.logError(r, err.Error(), status, nil, userID)
			handleApiError(rw, err)
		}
	}
}

func writeJson(rw http.ResponseWriter, status int, v any) error {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)

	if v != nil {
		err := json.NewEncoder(rw).Encode(v)
		if err != nil {
			return models.NewAPIError("error writing json", http.StatusInternalServerError)
		}
	}

	return nil
}

func handleApiError(rw http.ResponseWriter, err error) {
	errorUUID := uuid.New()
	slog.Error(err.Error())
	rw.Header().Add("Content-Type", "application/json")

	switch v := err.(type) {
	case *models.APIError:
		{
			switch status := v.HttpStatus; status {
			case 400:
				handleBadRequest(rw, err)
				return
			case 401:
				handleUnauthorized(rw, err)
				return
			case 402:
				handlePaymentRequired(rw, err)
				return
			default:
				handleDefault(rw, err, v.HttpStatus, errorUUID)
				return
			}
		}
	default:
		handleDefault(rw, err, http.StatusInternalServerError, errorUUID)
		return
	}
}

func handleBadRequest(rw http.ResponseWriter, err error) {
	result := models.APIErrorResponse{
		HttpStatus: http.StatusBadRequest,
		Errors:     []string{err.Error()[5:]},
	}

	rw.WriteHeader(http.StatusBadRequest)
	err = json.NewEncoder(rw).Encode(result)
	if err != nil {
		slog.Error("writing json")
	}
}

func handleUnauthorized(rw http.ResponseWriter, err error) {
	result := models.APIErrorResponse{
		HttpStatus: http.StatusUnauthorized,
		Errors:     []string{"Unauthorized"},
	}

	rw.WriteHeader(http.StatusUnauthorized)
	err = json.NewEncoder(rw).Encode(result)
	if err != nil {
		slog.Error("writing json")
	}
}

func handlePaymentRequired(rw http.ResponseWriter, err error) {
	result := models.APIErrorResponse{
		HttpStatus: http.StatusPaymentRequired,
		Errors:     []string{err.Error()[5:]},
	}

	rw.WriteHeader(http.StatusPaymentRequired)
	err = json.NewEncoder(rw).Encode(result)
	if err != nil {
		slog.Error("writing json")
	}
}

func handleDefault(rw http.ResponseWriter, err error, status int, errorUUID uuid.UUID) {
	rw.WriteHeader(status)
	err = json.NewEncoder(rw).Encode(models.APIErrorWithTraceResponse{
		TraceID: errorUUID.String(),
		APIErrorResponse: models.APIErrorResponse{
			HttpStatus: status,
			Errors:     []string{"server error"},
		},
	})
	if err != nil {
		slog.Error("writing json")
	}
}

func (s *Server) logError(r *http.Request, message string, statusCode int, stackTrace *string, userID *uuid.UUID) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	requestBodyString := string(b[:])

	toInsert := models.LogEdit{
		Message:     message,
		ApiUrl:      r.URL.String(),
		StackTrace:  stackTrace,
		Level:       models.LogLevelError,
		Source:      models.LogSourceAPI,
		UserID:      userID,
		RequestBody: &requestBodyString,
		StatusCode:  int32(statusCode),
	}
	_, err = s.db.InsertLog(context.Background(), toInsert)
	if err != nil {
		slog.Error("Could not insert logs: " + err.Error())
	}
}
