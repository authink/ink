package ext

import (
	"net/http"

	"github.com/authink/ink.go/src/i18n"
	"github.com/gin-gonic/gin"
)

type Context gin.Context

func translateErrorMsg(c *gin.Context, err error) {
	if e, ok := err.(*ClientError); ok {
		e.Message = i18n.Translate(c, e.Code)
	}
}

func (c *Context) AbortWithClientError(err error) {
	ginContext := (*gin.Context)(c)
	translateErrorMsg(ginContext, err)

	ginContext.AbortWithStatusJSON(
		http.StatusBadRequest,
		err,
	)
}

func (c *Context) AbortWithUnauthorized(err error) {
	ginContext := (*gin.Context)(c)
	translateErrorMsg(ginContext, err)

	ginContext.AbortWithStatusJSON(
		http.StatusUnauthorized,
		err,
	)
}

func (c *Context) AbortWithForbidden(err error) {
	ginContext := (*gin.Context)(c)
	translateErrorMsg(ginContext, err)

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
