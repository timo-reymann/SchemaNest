//go:build prod

package ui

import (
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"net/http"
	"strings"
)
import "embed"
import "io/fs"

//go:embed all:nextjs/build
var nextDist embed.FS

func hasFile(fs fs.FS, path string) bool {
	f, err := fs.Open(path)
	if err == nil {
		f.Close()
		return true
	}
	return false
}

func Handler() func(writer http.ResponseWriter, request *http.Request) {
	nextJsFiles, err := fs.Sub(nextDist, "nextjs/build")
	if err != nil {
		panic(err)
	}
	httpFs := http.FS(nextJsFiles)
	fileServer := http.FileServer(httpFs)

	return func(w http.ResponseWriter, r *http.Request) {
		cleanPath := strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/")
		if !hasFile(nextJsFiles, cleanPath) {
			r.URL.Path = "/"
		}

		w.Header().Set("Last-Modified", buildinfo.BuildTimeRFC1123)
		fileServer.ServeHTTP(w, r)
	}
}
