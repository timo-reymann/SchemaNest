package api

import (
	"net/http"
)

type SchemaNestApi struct{}

func (s SchemaNestApi) ListJSONSchemas(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s SchemaNestApi) GetApiSchemaJsonSchemaIdentifier(w http.ResponseWriter, r *http.Request, identifier string) {
	//TODO implement me
	panic("implement me")
}

func (s SchemaNestApi) PostApiSchemaJsonSchemaIdentifier(w http.ResponseWriter, r *http.Request, identifier string) {
	//TODO implement me
	panic("implement me")
}

func (s SchemaNestApi) GetApiSchemaJsonSchemaIdentifierChannelChannel(w http.ResponseWriter, r *http.Request, identifier string, channel string) {
	//TODO implement me
	panic("implement me")
}

func (s SchemaNestApi) GetApiSchemaJsonSchemaIdentifierLatest(w http.ResponseWriter, r *http.Request, identifier string) {
	//TODO implement me
	panic("implement me")
}

func (s SchemaNestApi) GetApiSchemaJsonSchemaIdentifierVersionVersion(w http.ResponseWriter, r *http.Request, identifier string, version string) {
	//TODO implement me
	panic("implement me")
}

var _ ServerInterface = (*SchemaNestApi)(nil)

func NewSchemaNestApi() *SchemaNestApi {
	return &SchemaNestApi{}
}
