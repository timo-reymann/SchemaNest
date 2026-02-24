package schemanest_files

import (
	_ "embed"
)

//go:embed openapi.yml
var OpenapiSpec []byte

//go:embed NOTICE
var Notice []byte

//go:embed LICENSE
var License []byte
