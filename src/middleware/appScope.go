package middleware

import (
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/gin-gonic/gin"
)

func AppScope(appName string) gin.HandlerFunc {
	return ext.Handler(func(c *ext.Context) {
		if app := c.MustGet("app").(*model.App); app.Name == appName {
			c.Next()
		} else {
			c.AbortWithForbidden(ext.ERR_INVALID_APP)
		}
	})
}
