package middleware

import (
	"github.com/authink/ink.go/src/core"
	"github.com/gin-gonic/gin"
)

func SetupInk(ink *core.Ink) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ink", ink)
		c.Next()
	}
}
