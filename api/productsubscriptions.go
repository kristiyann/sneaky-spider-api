package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/models"
)

// GET .../api/subscriptions?top=10&skip=0&expand=product&filter.id=?
func (s *Server) GetSubscriptions(rw http.ResponseWriter, r *http.Request) error {
	topParam := r.URL.Query().Get("top")
	skipParam := r.URL.Query().Get("skip")
	expand := r.URL.Query().Get("expand")
	filterID := r.URL.Query().Get("filter.id")
	filterStatus := r.URL.Query().Get("filter.status")
	top := defaultTop
	skip := 0

	if skipParam != "" {
		parseResult, err := strconv.Atoi(skipParam)
		if err != nil {
			return models.NewInvalidParameters()
		}
		skip = parseResult
	}

	if topParam != "" {
		parseResult, err := strconv.Atoi(topParam)
		if err != nil {
			return models.NewInvalidParameters()
		}
		top = parseResult
	}

	pagination := models.PaginationParams{
		Top:  top,
		Skip: skip,
	}

	parsedFilterID := uuid.Nil
	if filterID != "" {
		parseResult, err := uuid.Parse(filterID)
		if err != nil {
			return models.NewInvalidParameters()
		}
		parsedFilterID = parseResult
	}

	if filterStatus != "" {
		if !models.IsValidProductSubscriptionStatus(filterStatus) {
			return models.NewInvalidParameters()
		}
	}

	filter := models.ProductSubscriptionFilter{
		ID:     parsedFilterID,
		Status: filterStatus,
	}

	result, err := s.logic.GetSubscriptions(r.Context(), pagination, filter, expand)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, result)
}

// POST .../api/subscriptions body: ProductSubscriptionEdit
func (s *Server) InsertSubscription(rw http.ResponseWriter, r *http.Request) error {
	if !contentTypeIsJson(rw, r) {
		return models.NewAPIError("expecting application/json, got something else", http.StatusUnsupportedMediaType)
	}

	var model models.ProductSubscriptionEdit

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	if model.Size == "" || model.ProductExternalID == "" || !models.IsValidVendor(model.Vendor) {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	ID, err := s.logic.InsertSubscription(r.Context(), model)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusCreated, ID)
}

// PUT .../api/subscriptions/update-status?id={subscription_id}&status={status}
func (s *Server) UpdateSubscriptionStatus(rw http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	if idParam == "" || status == "" {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	if !models.IsValidProductSubscriptionStatus(status) {
		return models.NewAPIError("invalid status value", http.StatusBadRequest)
	}

	ID, err := uuid.Parse(idParam)
	if err != nil {
		return models.NewAPIError("invalid uuid value", http.StatusBadRequest)
	}

	err = s.logic.UpdateSubscriptionStatus(r.Context(), ID, status)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, nil)
}

// DEL .../api/subscriptions?id={subscription_id}
func (s *Server) DeleteSubscription(rw http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	ID, err := uuid.Parse(idParam)
	if err != nil {
		return models.NewAPIError("invalid uuid", http.StatusBadRequest)
	}

	err = s.logic.DeleteSubscriptions(r.Context(), ID)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, nil)
}
