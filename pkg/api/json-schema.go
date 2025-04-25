package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/mapping"
	"io"
	"net/http"
)

func (s *SchemaNestApi) ListJSONSchemas(w http.ResponseWriter, r *http.Request) {
	entities, err := s.context.JsonSchemaRepository.List(r.Context())
	if err != nil {
		sendInternalErr(w, "Failed to list JSON schemas", err)
		return
	}

	models := mapping.MapEntitiesToModel(
		func(e *json_schema.JsonSchemaEntity) *JsonSchemaInfo {
			return &JsonSchemaInfo{
				Identifier: &e.Identifier,
			}
		},
		entities,
	)
	_ = sendJSON(w, models)
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifier(w http.ResponseWriter, r *http.Request, identifier string) {
	entities, err := s.context.JsonSchemaVersionRepository.ListForJsonSchema(r.Context(), identifier)
	if err != nil {
		sendInternalErr(w, "Failed to list JSON schema versions for", identifier, err)
		return
	}
	models := mapping.MapEntitiesToModel(
		func(e *json_schema.JsonSchemaVersionEntity) *JsonSchemaVersion {
			version := fmt.Sprintf("%d.%d.%d", e.VersionMajor, e.VersionMinor, e.VersionPatch)
			return &JsonSchemaVersion{
				Version: &version,
			}
		},
		entities,
	)
	_ = sendJSON(w, models)
}

func (s *SchemaNestApi) PostApiSchemaJsonSchemaIdentifierVersionVersion(w http.ResponseWriter, r *http.Request, identifier string, version string) {
	v, err := semver.NewVersion(version)
	if err != nil || v.Metadata() != "" {
		sendError(w, http.StatusBadRequest, "Invalid version format. Only 'major.minor.patch' is supported.")
		return
	}

	raw, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	err = json.NewDecoder(bytes.NewBuffer(raw)).Decode(&map[string]any{})
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid JSON schema")
		return
	}

	err = s.context.JsonSchemaVersionRepository.Insert(r.Context(), &json_schema.JsonSchemaVersionEntity{
		Id:           nil,
		VersionMajor: v.Major(),
		VersionMinor: v.Minor(),
		VersionPatch: v.Patch(),
		Content:      string(raw),
		JsonSchemaId: 1,
	})
	if err != nil {
		sendError(w, http.StatusConflict, "Schema already exists")
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierChannelChannel(w http.ResponseWriter, r *http.Request, identifier string, channel string) {
	//TODO implement me
	panic("implement me")
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierLatest(w http.ResponseWriter, r *http.Request, identifier string) {
	//TODO implement me
	panic("implement me")
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierVersionVersion(w http.ResponseWriter, r *http.Request, identifier string, version string) {
	semver, err := semver.NewVersion(version)
	if err != nil || semver.Metadata() != "" {
		sendError(w, http.StatusBadRequest, "Invalid version format. Only 'major.minor.patch' is supported.")
		return
	}

	entity, err := s.context.JsonSchemaVersionRepository.GetForJsonSchemaAndVersion(r.Context(), identifier, semver.Major(), semver.Minor(), semver.Patch())
	if err != nil {
		sendInternalErr(w, "Failed to get JSON schema version for", identifier, err)
		return
	}

	if entity == nil {
		sendError(w, http.StatusNotFound, "Version not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(entity.Content))
}
