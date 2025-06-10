//go:build prod

package schemanest_files

import (
	"embed"
	_ "embed"
)

//go:embed all:ui/build
var UIFiles embed.FS
