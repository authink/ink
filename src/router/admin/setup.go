package admin

import (
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func SetupAdminGroup(rg *gin.RouterGroup, appName string) {
	gAdmin := rg.Group(
		"admin",
		inkstone.HandlerAdapter(middleware.AuthN),
		middleware.AppScope(appName),
	)
	gAdmin.GET("dashboard", dashboard)
	gAdmin.GET("apps", inkstone.HandlerAdapter(apps))
}
