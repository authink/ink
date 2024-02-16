package router

import (
	"github.com/authink/ink.go/src/core"
	_ "github.com/authink/ink.go/src/docs"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/router/admin"
	"github.com/authink/ink.go/src/router/token"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

//	@title			Ink API
//	@version		1.0
//	@description	This is ink server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	huoyijie
//	@contact.url	https://huoyijie.cn
//	@contact.email	yijie.huo@foxmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@Accept		json
//	@Produce	json

//	@securitydefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func SetupRouter(ink *core.Ink) (r *gin.Engine) {
	r = gin.Default()

	r.Use(
		middleware.SetupInk(ink),
		middleware.SetupI18n(ink),
	)

	gApi := r.Group("api/v1")
	token.SetupTokenGroup(gApi)
	admin.SetupAdminGroup(gApi, ink)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return
}
