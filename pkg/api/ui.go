package api

import (
	"encoding/json"
	"net/http"
)

func (s *SchemaNestApi) GetUiConfig(w http.ResponseWriter, r *http.Request) {
	cfg := s.context.Config
	res, _ := json.Marshal(UIConfig{
		ApiKeyAuthEnabled: cfg.EnableUploadAuthentication,
	})
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
