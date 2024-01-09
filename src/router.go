package main

import (
	"github.com/authink/ink.go/src/ext"
	"github.com/gin-gonic/gin"
)

func setupRouter(ink *Ink, r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		extContext := (*ext.Context)(c)
		extContext.AbortWithClientError(ext.ERR_CLI_BAD_EMAIL)
	})

	r.GET("/ping", func(c *gin.Context) {
		extContext := (*ext.Context)(c)
		extContext.AbortWithServerError(ext.ERR_SRV_DB_TIMEOUT)
	})
}
