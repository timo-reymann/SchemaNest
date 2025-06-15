package api

import (
	schemanestfiles "github.com/timo-reymann/SchemaNest"
	"net/http"
)

func (s *SchemaNestApi) GetApiSpecYml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/yaml")
	w.Header().Set("Content-Disposition", "inline;filename=api-spec.yml")
	_, _ = w.Write(schemanestfiles.OpenapiSpec)
}
