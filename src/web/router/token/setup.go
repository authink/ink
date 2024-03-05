package token

import (
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func SetupTokenGroup(rg *gin.RouterGroup) {
	gToken := rg.Group("token")
	gToken.POST("grant", web.HandlerAdapter(grant))
	gToken.POST("refresh", web.HandlerAdapter(refresh))
	gToken.POST("revoke", web.HandlerAdapter(revoke))
}
