package api

import (
	"github.com/timo-reymann/SchemaNest/pkg/encoding"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return encoding.WriteJSON(w, data)
}

func sendRawJSON(w http.ResponseWriter, data string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(data))
	return err
}

func sendInternalErr(w http.ResponseWriter, msg string, extra ...any) {
	slog.Info(msg, extra...)
	w.WriteHeader(http.StatusInternalServerError)
}

func sendError(w http.ResponseWriter, status int, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = encoding.WriteJSON(w, map[string]string{"error": error})
}
