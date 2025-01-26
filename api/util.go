package api

import (
	"net/http"
)

func (s *Server) authorize(f handlerFunc) func(rw http.ResponseWriter, r *http.Request) {
	return authorizationMiddleware(s.createHandler(f), s.supabaseClient, s.db)
}

func contentTypeIsJson(rw http.ResponseWriter, r *http.Request) bool {
	result := true
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		result = false
	}

	return result
}
