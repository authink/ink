package middleware

import (
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func AppScope(appName string) gin.HandlerFunc {
	return inkstone.HandlerAdapter(func(c *inkstone.Context) {
		if app := c.MustGet("app").(*model.App); app.Name == appName {
			c.Next()
		} else {
			c.AbortWithForbidden(errors.ERR_INVALID_APP)
		}
	})
}
