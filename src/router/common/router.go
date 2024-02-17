package common

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(ink *core.Ink) (router *gin.Engine, gApi *gin.RouterGroup) {
	router = gin.Default()

	router.Use(
		middleware.SetupInk(ink),
		middleware.SetupI18n(ink),
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gApi = router.Group(ink.Env.BasePath)
	return
}
