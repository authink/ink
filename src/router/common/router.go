package common

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(ink *core.Ink) (r *gin.Engine, gApi *gin.RouterGroup) {
	r = gin.Default()

	r.Use(
		middleware.SetupInk(ink),
		middleware.SetupI18n(ink),
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gApi = r.Group(ink.Env.BasePath)
	return
}
