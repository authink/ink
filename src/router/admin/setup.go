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
	gAdmin.GET("dashboard", inkstone.HandlerAdapter(dashboard))
	gAdmin.GET("apps", middleware.AuthZ("apps", "get"), inkstone.HandlerAdapter(apps))
	gAdmin.POST("apps", middleware.AuthZ("apps", "post"), inkstone.HandlerAdapter(addApp))
	gAdmin.PUT("apps/:id", middleware.AuthZ("apps", "put"), inkstone.HandlerAdapter(updateApp))
	gAdmin.GET("tokens", middleware.AuthZ("tokens", "get"), inkstone.HandlerAdapter(tokens))
	gAdmin.DELETE("tokens/:id", middleware.AuthZ("tokens", "delete"), inkstone.HandlerAdapter(deleteToken))
	gAdmin.GET("staffs", middleware.AuthZ("staffs", "get"), inkstone.HandlerAdapter(staffs))
	gAdmin.POST("staffs", middleware.AuthZ("staffs", "post"), inkstone.HandlerAdapter(addStaff))
	gAdmin.PUT("staffs/:id", middleware.AuthZ("staffs", "put"), inkstone.HandlerAdapter(updateStaff))
}
