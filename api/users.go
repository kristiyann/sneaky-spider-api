package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

var (
	supabaseRawUserDataKey_Name    = "name"
	supabaseRawUserDataKey_Picture = "picture"
)

// GET .../api/users
func (s *Server) GetCurrentUser(rw http.ResponseWriter, r *http.Request) error {
	session := util.GetSession(r.Context())
	result := models.UserSessionInfo{
		ID:                 session.ID,
		PublicID:           session.PublicID,
		Username:           session.Username,
		Email:              session.Email,
		ProfilePictureUrl:  session.Picture,
		Market:             session.Market,
		MembershipPlan:     util.GetUserMembershipPlanText(session.MembershipPlan, session.PlanEnds),
		DiscordWebhook:     session.DiscordWebhook,
		SnkrsGlobalEnabled: session.SnkrsGlobalEnabled,
	}

	return writeJson(rw, http.StatusOK, result)
}

// POST .../api/users/google/handle-signin body: models.TokensPayload
func (s *Server) HandleGoogleSignIn(rw http.ResponseWriter, r *http.Request) error {
	if !contentTypeIsJson(rw, r) {
		return models.NewAPIError("expecting application/json, got something else", http.StatusUnsupportedMediaType)
	}

	var model models.TokensPayload
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		return models.NewInvalidParameters()
	}

	if model.AccessToken == "" || model.RefreshToken == "" {
		return models.NewInvalidParameters()
	}

	authResult, err := s.supabaseClient.Auth.User(r.Context(), model.AccessToken)
	log.Println(authResult)
	if err != nil {
		return models.NewAPIError("retrieving supabase auth user: "+err.Error(), http.StatusInternalServerError)
	}

	authUserID, err := uuid.Parse(authResult.ID)
	if err != nil {
		handleApiError(rw, models.NewAPIError("problem decoding user ID", http.StatusInternalServerError))
	}

	err = s.logic.RegisterOrLoginFromProvider(r.Context(), authUserID, authResult.UserMetadata[supabaseRawUserDataKey_Name].(string), authResult.Email)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, nil)
}

// POST .../api/users/refresh-token body: models.TokensPayload
func (s *Server) RefreshToken(rw http.ResponseWriter, r *http.Request) error {
	if !contentTypeIsJson(rw, r) {
		return models.NewAPIError("expecting application/json, got something else", http.StatusUnsupportedMediaType)
	}

	var model models.TokensPayload
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		return models.NewInvalidParameters()
	}

	if model.RefreshToken == "" {
		return models.NewInvalidParameters()
	}

	refreshToken := model.RefreshToken
	accessToken := model.AccessToken

	authDetails, err := s.supabaseClient.Auth.RefreshUser(r.Context(), accessToken, refreshToken)
	if err != nil {
		return models.NewAPIError(fmt.Sprintf("couldnt refresh supabase user: %s", err.Error()), http.StatusInternalServerError)
	}

	result := models.RefreshTokenResponse{
		AccessToken:  authDetails.AccessToken,
		RefreshToken: authDetails.RefreshToken,
		TokenType:    authDetails.TokenType,
		ExpiresIn:    authDetails.ExpiresIn,
	}

	return writeJson(rw, http.StatusOK, result)
}

// POST .../api/users/signout
func (s *Server) SignOut(rw http.ResponseWriter, r *http.Request) error {
	supabaseToken := r.Context().Value(constants.ContextKeyUserSession).(models.UserSession).SupabaseToken

	err := s.supabaseClient.Auth.SignOut(r.Context(), supabaseToken)
	if err != nil {
		return models.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	return writeJson(rw, http.StatusOK, nil)
}

// PUT .../api/users/update-market?market={market}
func (s *Server) UpdateUserMarket(rw http.ResponseWriter, r *http.Request) error {
	market := r.URL.Query().Get("market")

	if market == "" {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	err := s.logic.UpdateUserMarket(r.Context(), market)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, nil)
}

// PUT .../api/users/monitors/snkrs?enabled={true/false}
func (s *Server) SetSnkrsGlobalEnabled(rw http.ResponseWriter, r *http.Request) error {
	enabled := r.URL.Query().Get("enabled")

	if enabled == "" {
		return models.NewInvalidParameters()
	}

	b, err := strconv.ParseBool(enabled)
	if err != nil {
		return models.NewAPIError("Could not convert string to boolean: "+err.Error(), http.StatusInternalServerError)
	}

	err = s.logic.SetSnkrsGlobalEnabled(r.Context(), b)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, nil)
}

// DEL .../api/users
func (s *Server) DeleteUserInfo(rw http.ResponseWriter, r *http.Request) error {
	err := s.logic.DeleteUser(r.Context())
	if err != nil {
		return err
	}

	supabaseToken := r.Context().Value(constants.ContextKeyUserSession).(models.UserSession).SupabaseToken

	err = s.supabaseClient.Auth.SignOut(r.Context(), supabaseToken)
	if err != nil {
		return models.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	return writeJson(rw, http.StatusOK, nil)
}
