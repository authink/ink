package middleware

import (
	"github.com/authink/ink.go/src/errs"
	"github.com/authink/ink.go/src/models"
	"github.com/authink/inkstone/web"
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
