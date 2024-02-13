package admin

import (
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/service"
	"github.com/gin-gonic/gin"
)

func SetupAdminGroup(rg *gin.RouterGroup) {
	gAdmin := rg.Group("admin", middleware.AuthN, middleware.AppScope(service.APP_ADMIN_DEV))
	gAdmin.GET("dashboard", dashboard)
}
