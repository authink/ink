package router

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/router/token"
	"github.com/gin-gonic/gin"
)

func SetupRouter(ink *core.Ink) (r *gin.Engine) {
	r = gin.Default()
	token.SetupTokenGroup(r)
	return
}
