package token

import (
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/inkstone"
)

// revoke godoc
//
//	@Summary		Revoke token
//	@Description	Revoke token
//	@Tags			token
//	@Router			/token/revoke [post]
//	@Param			refreshReq	body		refreshReq	true	"request body"
//	@Success		200			{string}	empty
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func revoke(c *inkstone.Context) {
	req := new(refreshReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	if _, ok := checkRefreshToken(c, req.RefreshToken); !ok {
		return
	}

	c.Empty()
}
