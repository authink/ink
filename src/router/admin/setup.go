package admin

import (
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAdminGroup(rg *gin.RouterGroup, appName string) {
	gAdmin := rg.Group(
		"admin",
		ext.Handler(middleware.AuthN),
		middleware.AppScope(appName),
	)
	gAdmin.GET("dashboard", dashboard)
	gAdmin.GET("apps", ext.Handler(apps))
}
