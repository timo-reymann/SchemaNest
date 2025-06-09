package ui

import (
	"github.com/timo-reymann/SchemaNest/pkg/api"
	"net/http"
	"strings"
)

func handleAPIRoute(w http.ResponseWriter, r *http.Request) bool {
	if strings.HasPrefix(r.URL.Path, "/api") {
		api.SendError(w, http.StatusNotFound, "unknown route")
		return true
	}
	return false
}
