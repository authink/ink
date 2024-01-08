package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter(ink *Ink, r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong pong pong",
		})
	})
}
