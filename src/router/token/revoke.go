package token

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
)

// revoke godoc
//
//	@Summary		Revoke token
//	@Description	Revoke token
//	@Tags			token
//	@Router			/token/revoke [post]
//	@Param			refreshReq	body		refreshReq	true	"request body"
//	@Success		200			{string}	empty
//	@Failure		400			{object}	ext.ClientError
//	@Failure		500			{string}	empty
func revoke(c *ext.Context) {
	req := &refreshReq{}
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
