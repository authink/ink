package middleware

import (
	"github.com/authink/ink.go/src/core"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func SetupI18n(ink *core.Ink) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Query("lang")
		accept := c.GetHeader("Accept-Language")
		localizer := i18n.NewLocalizer(ink.Bundle, lang, accept)

		c.Set("localizer", localizer)
		c.Next()
	}
}
