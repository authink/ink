package admin

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAdminGroup(rg *gin.RouterGroup, ink *core.Ink) {
	gAdmin := rg.Group(
		"admin",
		ext.Handler(middleware.AuthN),
		middleware.AppScope(ink.Env.AppNameAdmin),
	)
	gAdmin.GET("dashboard", dashboard)
}
