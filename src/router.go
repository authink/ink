package main

import (
	"net/http"

	"github.com/authink/ink.go/src/err"
	"github.com/gin-gonic/gin"
)

func setupRouter(ink *Ink, r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			err.ERR_CLI_BAD_EMAIL,
		)
		// c.AbortWithError(
		// 	http.StatusInternalServerError, err.ERR_SRV_DB_TIMEOUT,
		// )
	})
}
