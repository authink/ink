package ext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context gin.Context

func translateErrorMsg(err error) {
	if e, ok := err.(*ClientError); ok {
		// todo: Get i18n message by this code
		e.Message = e.Code
	}
}

func (c *Context) AbortWithClientError(err error) {
	translateErrorMsg(err)

	ginContext := (*gin.Context)(c)
	ginContext.AbortWithStatusJSON(
		http.StatusBadRequest,
		err,
	)
}

func (c *Context) AbortWithUnauthorized(err error) {
	translateErrorMsg(err)

	ginContext := (*gin.Context)(c)
	ginContext.AbortWithStatusJSON(
		http.StatusUnauthorized,
		err,
	)
}

func (c *Context) AbortWithForbidden(err error) {
	translateErrorMsg(err)

	ginContext := (*gin.Context)(c)
	ginContext.AbortWithStatusJSON(
		http.StatusForbidden,
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
