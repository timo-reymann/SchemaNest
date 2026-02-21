package api

import (
	"net/http"

	schemanestfiles "github.com/timo-reymann/SchemaNest"
)

func (s *SchemaNestApi) GetLicense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "inline;filename=LICENSE")
	_, _ = w.Write(schemanestfiles.License)
}

func (s *SchemaNestApi) GetNotice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "inline;filename=NOTICE")
	_, _ = w.Write(schemanestfiles.Notice)
}
