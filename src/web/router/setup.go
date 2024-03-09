package router

import (
	_ "github.com/authink/ink/src/docs"
	"github.com/authink/ink/src/envs"
	"github.com/authink/ink/src/web/router/admin"
	"github.com/authink/ink/src/web/router/token"
	"github.com/gin-gonic/gin"
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

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func SetupRouter(apiGroup *gin.RouterGroup) {
	token.SetupTokenGroup(apiGroup)
	admin.SetupAdminGroup(apiGroup, envs.AppNameAdmin())
}
