package static

import (
	"embed"
)

//go:embed embyExternalUrl/embyWebAddExternalUrl/embyLaunchPotplayer.js
//go:embed emby-web-mod/actorPlus/actorPlus.js
//go:embed emby-web-mod/emby-swiper/emby-swiper.js
//go:embed emby-web-mod/emby-tab/emby-tab.js
//go:embed emby-web-mod/fanart_show/fanart_show.js
//go:embed emby-web-mod/itemSortForNewDevice/itemSortForNewDevice.js
//go:embed emby-web-mod/playbackRate/playbackRate.js
//go:embed emby-web-mod/trailer/trailer.js
var EmbeddedStaticAssets embed.FS
