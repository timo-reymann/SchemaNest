package schemanest_files

import (
	_ "embed"
)

//go:embed openapi.yml
var OpenapiSpec []byte
