package api

import (
	"net/http"
	"strconv"

	"github.com/kristiyann/af1-spider-web-app/logic"
	"github.com/kristiyann/af1-spider-web-app/models"
)

// GET .../api/snkrs/search?market={BG}&search_term={text}&skip={}&top={}
func (s *Server) SearchSnkrsProducts(rw http.ResponseWriter, r *http.Request) error {
	searchTerm := r.URL.Query().Get("search_term")
	market := r.URL.Query().Get("market")
	top := r.URL.Query().Get("top")
	skip := r.URL.Query().Get("skip")

	if searchTerm == "" {
		return models.NewInvalidParameters()
	}

	if market == "" {
		market = "US"
	}

	skipNum, err := strconv.Atoi(skip)
	if err != nil {
		return err
	}

	topNum, err := strconv.Atoi(top)
	if err != nil {
		return err
	}

	params := logic.GetProductParams{
		Market: market,
		Skip:   skipNum,
		Top:    topNum,
	}

	result, err := s.logic.GetSnkrsProducts(r.Context(), params)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, result)
}
