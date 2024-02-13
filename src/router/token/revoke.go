package token

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
)

func revoke(c *ext.Context) {
	req := &reqRefresh{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(ext.ERR_BAD_REQUEST)
		return
	}

	ink := c.MustGet("ink").(*core.Ink)

	if _, ok := checkRefreshToken(c, ink, req.RefreshToken); !ok {
		return
	}

	c.Status(http.StatusOK)
}
