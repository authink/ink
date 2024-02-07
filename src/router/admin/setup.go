package admin

import "github.com/gin-gonic/gin"

func SetupAdminGroup(rg *gin.RouterGroup) {
	gAdmin := rg.Group("admin")
	gAdmin.GET("dashboard", dashboard)
}