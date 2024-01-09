package ext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context gin.Context

func (c *Context) AbortWithClientError(err error) {
	if e, ok := err.(*ClientError); ok {
		// todo: Get i18n message by this code
		e.Message = e.Code
	}

	ginContext := (*gin.Context)(c)
	ginContext.AbortWithStatusJSON(
		http.StatusBadRequest,
		err,
	)
}

func (c *Context) AbortWithServerError(err error) {
	ginContext := (*gin.Context)(c)
	ginContext.AbortWithError(
		http.StatusInternalServerError,
		err,
	)
}
