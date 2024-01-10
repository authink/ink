package token

import (
	"github.com/gin-gonic/gin"
)

func SetupTokenGroup(r *gin.Engine) {
	gToken := r.Group("token")
	gToken.POST("grant", grant)
}
