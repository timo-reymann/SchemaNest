package api

import (
	"github.com/timo-reymann/SchemaNest/pkg/config"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
)

type SchemaNestApiContext struct {
	JsonSchemaRepository        json_schema.JsonSchemaRepository
	JsonSchemaVersionRepository json_schema.JsonSchemaVersionRepository
	Config                      *config.Config
}

type SchemaNestApi struct {
	context *SchemaNestApiContext
}

var _ ServerInterface = (*SchemaNestApi)(nil)

func NewSchemaNestApi(ctx *SchemaNestApiContext) *SchemaNestApi {
	return &SchemaNestApi{ctx}
}
