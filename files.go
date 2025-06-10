package schemanest_files

import (
	"embed"
	_ "embed"
)

//go:embed openapi.yml
var OpenapiSpec []byte

//go:embed all:ui/build
var UIFiles embed.FS
