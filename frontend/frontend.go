//go:build embed

package frontend

import (
	"embed"
)

//go:embed dist
var Files embed.FS
