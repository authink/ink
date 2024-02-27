package middleware

import (
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func AuthZ() gin.HandlerFunc {
	return inkstone.HandlerAdapter(func(c *inkstone.Context) {
		
	})
}