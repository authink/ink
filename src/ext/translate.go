package ext

import "github.com/nicksnyder/go-i18n/v2/i18n"

func Translate(c *Context, messageID string) string {
	localizer := c.MustGet("localizer").(*i18n.Localizer)
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
}
