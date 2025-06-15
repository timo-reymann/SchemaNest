package api

import (
	schemanestfiles "github.com/timo-reymann/SchemaNest"
	"github.com/timo-reymann/SchemaNest/pkg/config"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
	"net/http"
)

type SchemaNestApiContext struct {
	JsonSchemaRepository        json_schema.JsonSchemaRepository
	JsonSchemaVersionRepository json_schema.JsonSchemaVersionRepository
	Config                      *config.Config
}

type SchemaNestApi struct {
	context *SchemaNestApiContext
}

func (s *SchemaNestApi) GetApiSpecYml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/yaml")
	w.Header().Set("Content-Disposition", "inline;filename=api-spec.yml")
	_, _ = w.Write(schemanestfiles.OpenapiSpec)
}

var _ ServerInterface = (*SchemaNestApi)(nil)

func NewSchemaNestApi(ctx *SchemaNestApiContext) *SchemaNestApi {
	return &SchemaNestApi{ctx}
}
