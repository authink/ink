package router

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/router/token"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ink *core.Ink) (r *gin.Engine) {
	r = gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("ink", ink)
		c.Next()
	})

	token.SetupTokenGroup(r.Group("api"))
	return
}
