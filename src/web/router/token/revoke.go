package token

import (
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/stone/web"
)

// revoke godoc
//
//	@Summary		Revoke token
//	@Description	Revoke token
//	@Tags			token
//	@Router			/token/revoke [post]
//	@Param			refreshReq	body		refreshReq	true	"request body"
//	@Success		200			{string}	empty
//	@Failure		400			{object}	web.ClientError
//	@Failure		500			{string}	empty
func revoke(c *web.Context) {
	req := &refreshReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	if _, ok := checkRefreshToken(c, req.RefreshToken); !ok {
		return
	}

	c.Empty()
}
