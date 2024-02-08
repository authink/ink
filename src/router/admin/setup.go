package admin

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAdminGroup(rg *gin.RouterGroup) {
	gAdmin := rg.Group("admin", middleware.AuthN, middleware.AppScope(core.APP_ADMIN_DEV))
	gAdmin.GET("dashboard", dashboard)
}
