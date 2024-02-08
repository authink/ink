package middleware

import (
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/gin-gonic/gin"
)

func AppScope(appName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		extCtx := (*ext.Context)(c)
		if app := c.MustGet("app").(*model.App); app.Name == appName {
			c.Next()
		} else {
			extCtx.AbortWithForbidden(ext.ERR_INVALID_APP)
		}
	}
}
