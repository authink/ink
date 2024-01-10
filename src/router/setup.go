package router

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ink *core.Ink) (r *gin.Engine) {

	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		extContext := (*ext.Context)(c)
		extContext.AbortWithClientError(ext.ERR_CLI_BAD_EMAIL)
	})

	r.GET("/ping", func(c *gin.Context) {
		extContext := (*ext.Context)(c)
		extContext.AbortWithServerError(ext.ERR_SRV_DB_TIMEOUT)
	})

	return
}
