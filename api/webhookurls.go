package api

import "net/http"

// PUT .../api/webhook-urls?url={url}
func (s *Server) UpdateWebhookUrl(rw http.ResponseWriter, r *http.Request) error {
	discordWebhookUrl := r.URL.Query().Get("url")
	err := s.logic.UpdateWebhook(r.Context(), discordWebhookUrl)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, nil)
}
