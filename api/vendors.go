package api

import (
	"net/http"

	"github.com/kristiyann/af1-spider-web-app/models"
)

// GET .../api/vendors
func (s *Server) GetVendors(rw http.ResponseWriter, r *http.Request) error {
	result := []string{
		models.VendorNike,
		models.VendorSNKRS,
	}

	return writeJson(rw, http.StatusOK, result)
}
