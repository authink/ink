package middleware

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
)

func AppScope(appName string) gin.HandlerFunc {
	return web.HandlerAdapter(func(c *web.Context) {
		if app := c.MustGet("app").(*models.App); app.Name == appName {
			c.Next()
		} else {
			c.AbortWithForbidden(errs.ERR_INVALID_APP)
		}
	})
}
