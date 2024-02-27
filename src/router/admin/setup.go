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
		inkstone.HandlerAdapter(middleware.AuthZ()),
	)
	gAdmin.GET("dashboard", inkstone.HandlerAdapter(dashboard))
	gAdmin.GET("apps", inkstone.HandlerAdapter(apps))
	gAdmin.POST("apps", inkstone.HandlerAdapter(addApp))
	gAdmin.PUT("apps/:id", inkstone.HandlerAdapter(updateApp))
	gAdmin.GET("tokens", inkstone.HandlerAdapter(tokens))
	gAdmin.DELETE("tokens/:id", inkstone.HandlerAdapter(deleteToken))
	gAdmin.GET("staffs", inkstone.HandlerAdapter(staffs))
	gAdmin.POST("staffs", inkstone.HandlerAdapter(addStaff))
	gAdmin.PUT("staffs/:id", inkstone.HandlerAdapter(updateStaff))
}
