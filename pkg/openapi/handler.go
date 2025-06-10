package openapi

import (
	schemanestfiles "github.com/timo-reymann/SchemaNest"
	"net/http"
)

func SpecHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/api-spec.yml" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/yaml")
	w.Header().Set("Content-Disposition", "inline;filename=api-spec.yml")
	_, _ = w.Write(schemanestfiles.OpenapiSpec)
}
