package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/timo-reymann/SchemaNest/pkg/channel"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/mapping"
	"io"
	"net/http"
	"strings"
)

const errMsgVersionNotFound = "version not found"

func (s *SchemaNestApi) ListJSONSchemas(w http.ResponseWriter, r *http.Request) {
	entities, err := s.context.JsonSchemaRepository.List(r.Context())
	if err != nil {
		sendInternalErr(w, "Failed to list JSON schemas", err)
		return
	}

	models := mapping.MapEntitiesToModel(
		func(e *json_schema.JsonSchemaEntityWithBasicInfo) *JsonSchemaInfo {
			return &JsonSchemaInfo{
				Identifier:  e.Identifier,
				Description: e.Description,
				LatestVersion: VersionParts{
					Major: e.LatestVersion.Major,
					Minor: e.LatestVersion.Minor,
					Patch: e.LatestVersion.Patch,
				},
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
	_, versions := mapping.MapEntitiesToModelWithValues(
		func(e *json_schema.JsonSchemaVersionEntity) *JsonSchemaVersion {
			version := fmt.Sprintf("%d.%d.%d", e.VersionMajor, e.VersionMinor, e.VersionPatch)
			return &JsonSchemaVersion{
				Version: version,
			}
		},
		entities,
	)

	latest, err := s.context.JsonSchemaVersionRepository.GetLatestVersion(r.Context(), identifier)
	if err != nil {
		sendInternalErr(w, "Failed to get latest JSON schema version for", identifier, err)
		return
	}

	details := JsonSchemaDetails{
		Versions:    versions,
		Description: latest.Description,
	}
	_ = sendJSON(w, details)
}

func (s *SchemaNestApi) PostApiSchemaJsonSchemaIdentifierVersionVersion(w http.ResponseWriter, r *http.Request, identifier string, version string) {
	if s.context.Config.EnableUploadAuthentication && !s.isAuthenticated(w, r, identifier) {
		return
	}

	v, err := semver.NewVersion(version)
	if err != nil || v.Metadata() != "" {
		SendError(w, http.StatusBadRequest, "invalid version format. Only 'major.minor.patch' is supported.")
		return
	}

	raw, err := io.ReadAll(r.Body)
	if err != nil {
		SendError(w, http.StatusBadRequest, "failed to read request body")
		return
	}

	decoded := map[string]any{}
	err = json.NewDecoder(bytes.NewBuffer(raw)).Decode(&decoded)
	if err != nil {
		SendError(w, http.StatusBadRequest, "invalid JSON schema: "+err.Error())
		return
	}

	var description *string
	if val, ok := decoded["description"]; ok {
		valStr := val.(string)
		description = &valStr
	}

	schema, err := s.context.JsonSchemaRepository.Get(r.Context(), identifier)
	if err != nil {
		err = s.context.JsonSchemaRepository.Insert(r.Context(), &json_schema.JsonSchemaEntity{
			Identifier: identifier,
		})
		schema, err = s.context.JsonSchemaRepository.Get(r.Context(), identifier)
	}

	if err != nil || schema == nil {
		SendError(w, http.StatusInternalServerError, "failed to get or create json schema entry")
		return
	}

	err = s.context.JsonSchemaVersionRepository.Insert(r.Context(), &json_schema.JsonSchemaVersionEntity{
		Id:           nil,
		Description:  description,
		VersionMajor: v.Major(),
		VersionMinor: v.Minor(),
		VersionPatch: v.Patch(),
		Content:      string(raw),
		JsonSchemaId: *schema.Id,
	})
	if err != nil {
		SendError(w, http.StatusConflict, "schema already exists")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *SchemaNestApi) isAuthenticated(w http.ResponseWriter, r *http.Request, identifier string) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		SendError(w, http.StatusUnauthorized, "authorization header not set")
		return false
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		SendError(w, http.StatusBadRequest, "invalid authorization header")
		return false
	}

	apiKey, found := s.context.Config.LookupApiKey(headerParts[1])
	if !found {
		SendError(w, http.StatusForbidden, "invalid api key")
		return false
	}

	if !apiKey.IsUsableForSchemaIdentifier(identifier) {
		SendError(w, http.StatusForbidden, fmt.Sprintf("the api key you provided can not be used to upload JSON schemas with the identifier '%s'", identifier))
		return false
	}
	return true
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierChannelChannel(w http.ResponseWriter, r *http.Request, identifier string, channelIdentifier string) {
	parsedChannel, err := channel.Parse(channelIdentifier)
	if err != nil {
		SendError(w, http.StatusBadRequest, "invalid channel format: "+err.Error())
		return
	}

	var entity *json_schema.JsonSchemaVersionEntity
	if parsedChannel.Minor == "" {
		entity, err = s.context.JsonSchemaVersionRepository.GetForLatestMajorVersion(r.Context(), identifier, int64(parsedChannel.MajorVersion()))
	} else {
		entity, err = s.context.JsonSchemaVersionRepository.GetForLatestMinorVersion(r.Context(), identifier, int64(parsedChannel.MajorVersion()), int64(parsedChannel.MinorVersion()))
	}

	if entity == nil {
		SendError(w, http.StatusNotFound, errMsgVersionNotFound)
		return
	}

	_ = sendRawJSON(w, entity.Content)
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierLatest(w http.ResponseWriter, r *http.Request, identifier string) {
	entity, err := s.context.JsonSchemaVersionRepository.GetLatestVersion(r.Context(), identifier)
	if err != nil {
		sendInternalErr(w, "Failed to get latest JSON schema version for", identifier, err)
		return
	}

	if entity == nil {
		SendError(w, http.StatusNotFound, errMsgVersionNotFound)
		return
	}

	_ = sendRawJSON(w, entity.Content)
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierLatestVersion(w http.ResponseWriter, r *http.Request, identifier string) {
	entity, err := s.context.JsonSchemaVersionRepository.GetLatestVersion(r.Context(), identifier)
	if entity == nil || err != nil {
		SendError(w, http.StatusNotFound, errMsgVersionNotFound)
		return
	}

	_ = sendJSON(w, VersionParts{
		Major: int(entity.VersionMajor),
		Minor: int(entity.VersionMinor),
		Patch: int(entity.VersionPatch),
	})
}

func (s *SchemaNestApi) GetApiSchemaJsonSchemaIdentifierVersionVersion(w http.ResponseWriter, r *http.Request, identifier string, version string) {
	semver, err := semver.NewVersion(version)
	if err != nil || semver.Metadata() != "" {
		SendError(w, http.StatusBadRequest, "invalid version format. Only 'major.minor.patch' is supported.")
		return
	}

	entity, err := s.context.JsonSchemaVersionRepository.GetForJsonSchemaAndVersion(r.Context(), identifier, semver.Major(), semver.Minor(), semver.Patch())
	if err != nil {
		sendInternalErr(w, "failed to get JSON schema version for", identifier, err)
		return
	}

	if entity == nil {
		SendError(w, http.StatusNotFound, errMsgVersionNotFound)
		return
	}

	_ = sendRawJSON(w, entity.Content)
}
