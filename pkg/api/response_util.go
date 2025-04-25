package api

import (
	"github.com/timo-reymann/SchemaNest/pkg/encoding"
	"log/slog"
	"net/http"
)

func SendJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return encoding.WriteJSON(w, data)
}

func SendInternalErr(w http.ResponseWriter, msg string, extra ...any) {
	slog.Info(msg, extra...)
	w.WriteHeader(http.StatusInternalServerError)
}

func SendError(w http.ResponseWriter, status int, error string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_ = encoding.WriteJSON(w, map[string]string{"error": error})
}
