package ext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func translateErrorMsg(c *Context, err error) {
	if e, ok := err.(*ClientError); ok {
		e.Message = Translate(c, e.Code)
	}
}

func (c *Context) AbortWithClientError(err error) {
	translateErrorMsg(c, err)
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		err,
	)
}

func (c *Context) AbortWithUnauthorized(err error) {
	translateErrorMsg(c, err)
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		err,
	)
}

func (c *Context) AbortWithForbidden(err error) {
	translateErrorMsg(c, err)
	c.AbortWithStatusJSON(
		http.StatusForbidden,
		err,
	)
}

func (c *Context) AbortWithServerError(err error) {
	c.AbortWithError(
		http.StatusInternalServerError,
		err,
	)
}
