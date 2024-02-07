package token

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/gin-gonic/gin"
)

func revoke(c *gin.Context) {
	extCtx := (*ext.Context)(c)

	req := &reqRefresh{}
	if err := c.ShouldBindJSON(req); err != nil {
		extCtx.AbortWithClientError(ext.ERR_BAD_REQUEST)
		return
	}

	ink := c.MustGet("ink").(*core.Ink)

	if _, ok := checkRefreshToken(extCtx, ink, req.RefreshToken); !ok {
		return
	}

	c.Status(http.StatusOK)
}
