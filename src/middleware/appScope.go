package middleware

import (
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func AppScope(appName string) gin.HandlerFunc {
	return web.HandlerAdapter(func(c *web.Context) {
		if app := c.MustGet("app").(*model.App); app.Name == appName {
			c.Next()
		} else {
			c.AbortWithForbidden(errors.ERR_INVALID_APP)
		}
	})
}
