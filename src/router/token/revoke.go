package token

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/gin-gonic/gin"
)

func revoke(c *gin.Context) {
	req := &reqRefresh{}
	if err := c.BindJSON(req); err != nil {
		return
	}

	extCtx := (*ext.Context)(c)
	ink := c.MustGet("ink").(*core.Ink)

	if _, ok := checkRefreshToken(extCtx, ink, req.RefreshToken); !ok {
		return
	}

	c.Status(http.StatusOK)
}
