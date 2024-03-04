package i18n

import (
	"embed"
)

//go:embed locales/*.toml
var Locales embed.FS
