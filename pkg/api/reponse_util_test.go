package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     any
		expected string
		wantErr  bool
	}{
		{"Valid JSON", map[string]string{"key": "value"}, `{"key":"value"}` + "\n", false},
		{"Nil Data", nil, "null\n", false},
		{"Invalid Data", make(chan int), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			err := sendJSON(rec, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && rec.Body.String() != tt.expected {
				t.Errorf("sendJSON() = %q, want %q", rec.Body.String(), tt.expected)
			}
			if !tt.wantErr && rec.Header().Get("Content-Type") != "application/json" {
				t.Errorf("sendJSON() Content-Type = %q, want %q", rec.Header().Get("Content-Type"), "application/json")
			}
		})
	}
}

func TestSendInternalErr(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{"Internal Error", "Something went wrong"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			sendInternalErr(rec, tt.msg)
			if rec.Code != http.StatusInternalServerError {
				t.Errorf("sendInternalErr() status = %d, want %d", rec.Code, http.StatusInternalServerError)
			}
		})
	}
}

func TestSendError(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		errorMsg string
		expected string
	}{
		{"Bad Request", http.StatusBadRequest, "Invalid input", `{"error":"Invalid input"}` + "\n"},
		{"Not Found", http.StatusNotFound, "Resource not found", `{"error":"Resource not found"}` + "\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			sendError(rec, tt.status, tt.errorMsg)
			if rec.Code != tt.status {
				t.Errorf("sendError() status = %d, want %d", rec.Code, tt.status)
			}
			if rec.Body.String() != tt.expected {
				t.Errorf("sendError() = %q, want %q", rec.Body.String(), tt.expected)
			}
			if rec.Header().Get("Content-Type") != "application/json" {
				t.Errorf("sendError() Content-Type = %q, want %q", rec.Header().Get("Content-Type"), "application/json")
			}
		})
	}
}
