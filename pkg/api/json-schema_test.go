package api

import (
	"bytes"
	"encoding/json"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListJSONSchemas(t *testing.T) {
	tests := []struct {
		name        string
		mockSchemas []*json_schema.JsonSchemaEntityWithBasicInfo
		expected    int
	}{
		{
			name: "Success",
			mockSchemas: []*json_schema.JsonSchemaEntityWithBasicInfo{
				{
					JsonSchemaEntity: json_schema.JsonSchemaEntity{
						Identifier: "schema1",
					},
				},
				{
					JsonSchemaEntity: json_schema.JsonSchemaEntity{
						Identifier: "schema2",
					},
				},
			},
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContext := createTestContext()
			mockRepo := mockContext.JsonSchemaRepository.(*MockJsonSchemaRepository)
			mockRepo.Schemas = tt.mockSchemas

			api := NewSchemaNestApi(&mockContext)
			req := httptest.NewRequest(http.MethodGet, "/schemas", nil)
			rec := httptest.NewRecorder()

			api.ListJSONSchemas(rec, req)

			if rec.Code != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, rec.Code)
			}
			if tt.expected == http.StatusOK {
				var response []JsonSchemaInfo
				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t.Errorf("failed to decode response: %v", err)
				}
				if len(response) != len(tt.mockSchemas) {
					t.Errorf("expected %d schemas, got %d", len(tt.mockSchemas), len(response))
				}
			}
		})
	}
}

func TestGetApiSchemaJsonSchemaIdentifier(t *testing.T) {
	tests := []struct {
		name         string
		mockVersions []*json_schema.JsonSchemaVersionEntity
		expected     int
	}{
		{
			name: "Success",
			mockVersions: []*json_schema.JsonSchemaVersionEntity{
				{JsonSchemaId: 1, VersionMajor: 1, VersionMinor: 0, VersionPatch: 0},
				{JsonSchemaId: 1, VersionMajor: 2, VersionMinor: 0, VersionPatch: 0},
			},
			expected: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContext := createTestContext()
			mockRepo := mockContext.JsonSchemaVersionRepository.(*MockJsonSchemaVersionRepository)
			mockRepo.Versions = tt.mockVersions

			api := NewSchemaNestApi(&mockContext)
			req := httptest.NewRequest(http.MethodGet, "/schemas/schema1", nil)
			rec := httptest.NewRecorder()

			api.GetApiSchemaJsonSchemaIdentifier(rec, req, "schema1")

			if rec.Code != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, rec.Code)
			}
			if tt.expected == http.StatusOK {
				response := JsonSchemaDetails{}
				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t.Errorf("failed to decode response: %v", err)
				}
				if len(response.Versions) != len(tt.mockVersions) {
					t.Errorf("expected %d versions, got %d", len(tt.mockVersions), len(response.Versions))
				}
			}
		})
	}
}

func TestPostApiSchemaJsonSchemaIdentifierVersionVersion(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		version  string
		expected int
	}{
		{
			name:     "Valid Schema",
			body:     `{"type": "object"}`,
			version:  "1.0.0",
			expected: http.StatusCreated,
		},
		{
			name:     "Invalid Version",
			body:     `{"type": "object"}`,
			version:  "1.a",
			expected: http.StatusBadRequest,
		},
		{
			name:     "Invalid JSON",
			body:     `{"type": `,
			version:  "1.0.0",
			expected: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContext := createTestContext()
			mockRepo := mockContext.JsonSchemaRepository.(*MockJsonSchemaRepository)
			var id int64 = 1
			mockRepo.Schemas = []*json_schema.JsonSchemaEntityWithBasicInfo{
				{
					JsonSchemaEntity: json_schema.JsonSchemaEntity{Identifier: "schema1", Id: &id},
				},
			}
			api := NewSchemaNestApi(&mockContext)

			req := httptest.NewRequest(http.MethodPost, "/schemas/schema1/"+tt.version, bytes.NewBufferString(tt.body))
			rec := httptest.NewRecorder()

			api.PostApiSchemaJsonSchemaIdentifierVersionVersion(rec, req, "schema1", tt.version)

			if rec.Code != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, rec.Code)
			}
		})
	}
}

func TestGetApiSchemaJsonSchemaIdentifierVersionVersion(t *testing.T) {
	tests := []struct {
		name       string
		version    string
		mockEntity *json_schema.JsonSchemaVersionEntity
		mockError  error
		expected   int
	}{
		{
			name:    "Success",
			version: "1.0.0",
			mockEntity: &json_schema.JsonSchemaVersionEntity{
				JsonSchemaId: 1, VersionMajor: 1, VersionMinor: 0, VersionPatch: 0, Content: `{"type": "object"}`,
			},
			expected: http.StatusOK,
		},
		{
			name:     "Not Found",
			version:  "2.0.0",
			expected: http.StatusNotFound,
		},
		{
			name:     "Invalid Version",
			version:  "1.a",
			expected: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContext := createTestContext()
			mockRepo := mockContext.JsonSchemaVersionRepository.(*MockJsonSchemaVersionRepository)
			if tt.mockEntity != nil {
				mockRepo.Versions = []*json_schema.JsonSchemaVersionEntity{tt.mockEntity}
			}

			api := NewSchemaNestApi(&mockContext)
			req := httptest.NewRequest(http.MethodGet, "/schemas/schema1/"+tt.version, nil)
			rec := httptest.NewRecorder()

			api.GetApiSchemaJsonSchemaIdentifierVersionVersion(rec, req, "schema1", tt.version)

			if rec.Code != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, rec.Code)
			}
			if tt.expected == http.StatusOK {
				if rec.Body.String() != tt.mockEntity.Content {
					t.Errorf("expected body %s, got %s", tt.mockEntity.Content, rec.Body.String())
				}
			}
		})
	}
}

func TestGetApiSchemaJsonSchemaIdentifierChannelChannel(t *testing.T) {
	tests := []struct {
		name           string
		channel        string
		mockEntity     *json_schema.JsonSchemaVersionEntity
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Valid channel with major version",
			channel: "1.x",
			mockEntity: &json_schema.JsonSchemaVersionEntity{
				JsonSchemaId: 1, VersionMajor: 1, VersionMinor: 0, VersionPatch: 0, Content: `{"type": "object"}`,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"type": "object"}`,
		},
		{
			name:    "Valid channel with major and minor version",
			channel: "1.0.x",
			mockEntity: &json_schema.JsonSchemaVersionEntity{
				JsonSchemaId: 1, VersionMajor: 1, VersionMinor: 0, VersionPatch: 0, Content: `{"type": "object"}`,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"type": "object"}`,
		},
		{
			name:           "Invalid channel format",
			channel:        "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid channel with single digit",
			channel:        "2",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContext := createTestContext()
			mockRepo := mockContext.JsonSchemaVersionRepository.(*MockJsonSchemaVersionRepository)
			if tt.mockEntity != nil {
				mockRepo.Versions = []*json_schema.JsonSchemaVersionEntity{tt.mockEntity}
			}

			api := NewSchemaNestApi(&mockContext)
			req := httptest.NewRequest(http.MethodGet, "/schemas/schema1/channel/"+tt.channel, nil)
			rec := httptest.NewRecorder()

			api.GetApiSchemaJsonSchemaIdentifierChannelChannel(rec, req, "schema1", tt.channel)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}
			if tt.expectedStatus == http.StatusOK && rec.Body.String() != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, rec.Body.String())
			}
		})
	}
}

func TestGetApiSchemaJsonSchemaIdentifierLatest(t *testing.T) {
	tests := []struct {
		name       string
		identifier string
		mockEntity *json_schema.JsonSchemaVersionEntity
		expected   int
		body       string
	}{
		{
			name:       "Valid Identifier",
			identifier: "test-schema",
			mockEntity: &json_schema.JsonSchemaVersionEntity{
				JsonSchemaId: 1, VersionMajor: 1, VersionMinor: 0, VersionPatch: 0, Content: `{"type": "object"}`,
			},
			expected: http.StatusOK,
			body:     `{"type": "object"}`,
		},
		{
			name:       "Invalid Identifier",
			identifier: "invalid-schema",
			mockEntity: nil,
			expected:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContext := createTestContext()
			mockRepo := mockContext.JsonSchemaVersionRepository.(*MockJsonSchemaVersionRepository)
			if tt.mockEntity != nil {
				mockRepo.Versions = []*json_schema.JsonSchemaVersionEntity{tt.mockEntity}
			}

			api := NewSchemaNestApi(&mockContext)
			req := httptest.NewRequest(http.MethodGet, "/schemas/"+tt.identifier+"/latest", nil)
			rec := httptest.NewRecorder()

			api.GetApiSchemaJsonSchemaIdentifierLatest(rec, req, tt.identifier)

			if rec.Code != tt.expected {
				t.Errorf("expected status %d, got %d", tt.expected, rec.Code)
			}
			if tt.expected == http.StatusOK && rec.Body.String() != tt.body {
				t.Errorf("expected body %s, got %s", tt.body, rec.Body.String())
			}
		})
	}
}
