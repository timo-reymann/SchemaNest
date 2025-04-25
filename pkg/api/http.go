package api

import (
	"net/http"
)

func NewServeMux(ctx *SchemaNestApiContext) (*http.ServeMux, error) {
	swagger, err := GetSwagger()
	if err != nil {
		return nil, err
	}
	swagger.Servers = nil

	r := http.NewServeMux()

	schemaNestApi := NewSchemaNestApi(ctx)
	HandlerFromMux(schemaNestApi, r)
	return r, nil
}
