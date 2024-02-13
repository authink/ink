package router

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/router/admin"
	"github.com/authink/ink.go/src/router/token"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ink *core.Ink) (r *gin.Engine) {
	r = gin.Default()

	r.Use(middleware.SetupInk(ink))
	r.Use(middleware.SetupI18n(ink))

	gApi := r.Group("api")
	token.SetupTokenGroup(gApi)
	admin.SetupAdminGroup(gApi)
	return
}
