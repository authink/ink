package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func setupAppGroup(gAdmin *gin.RouterGroup) {
	gApps := gAdmin.Group(authz.ResourceApp)
	gApps.Use(middleware.AuthZ(authz.ResourceApp))
	gApps.GET("", inkstone.HandlerAdapter(apps))
	gApps.POST("", inkstone.HandlerAdapter(addApp))
	gApps.PUT(":id", inkstone.HandlerAdapter(updateApp))
}

func setupTokenGroup(gAdmin *gin.RouterGroup) {
	gTokens := gAdmin.Group(authz.ResourceToken)
	gTokens.Use(middleware.AuthZ(authz.ResourceToken))
	gTokens.GET("", inkstone.HandlerAdapter(tokens))
	gTokens.DELETE(":id", inkstone.HandlerAdapter(deleteToken))
}

func setupStaffGroup(gAdmin *gin.RouterGroup) {
	gStaffs := gAdmin.Group(authz.ResourceStaff)
	gStaffs.Use(middleware.AuthZ(authz.ResourceStaff))
	gStaffs.GET("", inkstone.HandlerAdapter(staffs))
	gStaffs.POST("", inkstone.HandlerAdapter(addStaff))
	gStaffs.PUT(":id", inkstone.HandlerAdapter(updateStaff))
}

func SetupAdminGroup(rg *gin.RouterGroup, appName string) {
	gAdmin := rg.Group("admin")
	gAdmin.Use(
		inkstone.HandlerAdapter(middleware.AuthN), middleware.AppScope(appName),
	)
	gAdmin.GET("dashboard", inkstone.HandlerAdapter(dashboard))

	setupAppGroup(gAdmin)

	setupTokenGroup(gAdmin)

	setupStaffGroup(gAdmin)
}
