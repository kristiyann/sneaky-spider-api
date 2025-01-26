package api

import (
	"net/http"

	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
)

// GET .../api/nike?market={BG}&product_id={external_id}
func (s *Server) GetNikeProduct(rw http.ResponseWriter, r *http.Request) error {
	productId := r.URL.Query().Get("product_id")
	market := r.URL.Query().Get("market")

	if productId == "" {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	if market == "" {
		market = "US"
	}

	params := logic.GetProductParams{
		Identifier: productId,
		Market:     market,
	}

	result, err := s.logic.GetNikeProduct(r.Context(), params)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, result)
}

// GET .../api/nike/search?market={BG}&search_term={text}
func (s *Server) SearchNikeProducts(rw http.ResponseWriter, r *http.Request) error {
	searchTerm := r.URL.Query().Get("search_term")
	market := r.URL.Query().Get("market")

	if searchTerm == "" {
		return models.NewAPIError("invalid parameters", http.StatusBadRequest)
	}

	if market == "" {
		market = "US"
	}

	params := logic.SearchProductParams{
		SearchTerm: searchTerm,
		Market:     market,
	}

	result, err := s.logic.SearchNikeProducts(r.Context(), params)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, result)
}
