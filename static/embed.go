package static

import (
	"embed"
)

//go:embed embed.go
var EmbeddedStaticAssets embed.FS
