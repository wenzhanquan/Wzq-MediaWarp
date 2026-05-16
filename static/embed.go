package static

import (
	"embed"
)

//go:embed embyExternalUrl/embyWebAddExternalUrl/embyLaunchPotplayer.js
var EmbeddedStaticAssets embed.FS
