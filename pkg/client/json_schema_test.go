package client

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestClient_UploadJsonSchema(t *testing.T) {
	// Create a temporary JSON file for testing
	tempDir := t.TempDir()
	jsonFile := filepath.Join(tempDir, "test.json")
	if err := os.WriteFile(jsonFile, []byte(`{"test": "data"}`), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	tests := []struct {
		name       string
		identifier string
		version    string
		statusCode int
		response   string
		wantErr    bool
	}{
		{
			name:       "Successful upload",
			identifier: "test-schema",
			version:    "1.0.0",
			statusCode: http.StatusCreated,
			response:   `{}`,
			wantErr:    false,
		},
		{
			name:       "Conflict error",
			identifier: "existing-schema",
			version:    "1.0.0",
			statusCode: http.StatusConflict,
			response:   `{"error": "Schema already exists"}`,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request method and path
				if r.Method != http.MethodPost {
					t.Errorf("Expected POST request, got %s", r.Method)
				}
				expectedPath := "/api/schema/json-schema/" + tt.identifier + "/version/" + tt.version
				if r.URL.Path != expectedPath {
					t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
				}

				// Verify content type
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.response))
			}))
			defer server.Close()

			// Create client
			client, err := NewClient(server.URL, "")
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			// Test upload
			err = client.UploadJsonSchema(tt.identifier, tt.version, jsonFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadJsonSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
