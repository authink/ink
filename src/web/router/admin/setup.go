package admin

import (
	"github.com/authink/ink/src/web/middleware"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupAdminGroup(rg *gin.RouterGroup, appName string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(web.ValidationNotAllFieldsZero, &updateAppReq{}, &updateStaffReq{}, &updateGroupReq{})
	}

	gAdmin := rg.Group("admin")
	gAdmin.Use(
		web.HandlerAdapter(middleware.Authn), middleware.AppScope(appName),
	)
	setupDashboard(gAdmin)

	setupAppGroup(gAdmin)

	setupTokenGroup(gAdmin)

	setupStaffGroup(gAdmin)

	setupGroupGroup(gAdmin)

	setupGroupshipGroup(gAdmin)

	setupPermissionGroup(gAdmin)

	setupPolicyGroup(gAdmin)

	setupDeptGroup(gAdmin)

	setupLogGroup(gAdmin)
}
