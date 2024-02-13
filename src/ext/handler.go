package ext

import "github.com/gin-gonic/gin"

type HandlerFunc func(*Context)

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&Context{c})
	}
}
