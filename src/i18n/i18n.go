package i18n

import (
	"embed"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.toml
var locales embed.FS

func NewBundle() (bundle *i18n.Bundle) {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(&locales, "locales/en.toml")
	bundle.LoadMessageFileFS(&locales, "locales/zh-CN.toml")
	return
}
