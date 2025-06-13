//go:build prod

package ui

import (
	"github.com/gobwas/glob"
	schemanest_files "github.com/timo-reymann/SchemaNest"
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"net/http"
	"strings"
)
import "io/fs"

func handleStaticPage(r *RequestWithCleanPath) bool {
	if r.hasFile(r.cleanPath + ".html") {
		r.URL.Path = r.cleanPath + ".html"
		return true
	}
	return false
}

type DynamicRoute struct {
	Pattern glob.Glob
	RouteTo string
}

func handleDynamicRoutes(r *RequestWithCleanPath, mapping []*DynamicRoute) bool {
	for _, m := range mapping {
		if m.Pattern.Match(r.cleanPath) {
			r.URL.Path = m.RouteTo
			return true
		}
	}
	return false
}

type RequestWithCleanPath struct {
	*http.Request
	fs        fs.FS
	cleanPath string
}

func (rwp *RequestWithCleanPath) isIndex() bool {
	return rwp.cleanPath == ""
}

func (rwp *RequestWithCleanPath) isAsset() bool {
	return strings.HasPrefix(rwp.cleanPath, "_next")
}

func (rwp *RequestWithCleanPath) hasFile(path string) bool {
	f, err := rwp.fs.Open(path)
	if err == nil {
		f.Close()
		return true
	}
	return false
}

var dynamicRoutes = []*DynamicRoute{
	{
		glob.MustCompile("schemas/json-schema/*/*"),
		"/schemas/json-schema/[identifier]/[version].html",
	},
}

func CreateHandler() func(writer http.ResponseWriter, request *http.Request) {
	nextJsFiles, err := fs.Sub(schemanest_files.UIFiles, "ui/build")
	if err != nil {
		panic(err)
	}

	httpFs := http.FS(nextJsFiles)
	fileServer := http.FileServer(httpFs)

	return func(w http.ResponseWriter, r *http.Request) {
		if handleAPIRoute(w, r) {
			return
		}

		cleanPath := strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/")
		rwp := &RequestWithCleanPath{Request: r, cleanPath: cleanPath, fs: nextJsFiles}

		// Route index document and assets as is
		if !rwp.isIndex() && !rwp.isAsset() {
			// Try to handle other paths and fallback to 404 page
			if !handleStaticPage(rwp) && !handleDynamicRoutes(rwp, dynamicRoutes) {
				r.URL.Path = "/404.html"
			}
		}

		w.Header().Set("Last-Modified", buildinfo.BuildTimeRFC1123)
		fileServer.ServeHTTP(w, r)
	}
}
