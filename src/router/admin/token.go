package admin

import (
	"net/http"
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
)

type pageReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

type tokenRes struct {
	Id           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	AppId        uint32    `json:"appId"`
	AppName      string    `json:"appName"`
	AccountId    uint32    `json:"accountId"`
}

// tokens godoc
//
//	@Summary		Show tokens
//	@Description	Show tokens
//	@Tags			admin_token
//	@Router			/admin/tokens	[get]
//	@Security		ApiKeyAuth
//	@Param			offset	query		int	true	"offset"
//	@Param			limit	query		int	true	"limit"
//	@Success		200		{array}		tokenRes
//	@Failure		401		{object}	inkstone.ClientError
//	@Failure		403		{object}	inkstone.ClientError
//	@Failure		500		{string}	empty
func tokens(c *inkstone.Context) {
	appContext := c.App()

	req := &pageReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	tokens, err := orm.AuthToken(appContext).Pagination(req.Offset, req.Limit)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []tokenRes
	for i := range tokens {
		token := &tokens[i]
		res = append(res, tokenRes{
			int(token.Id),
			token.CreatedAt,
			token.UpdatedAt,
			token.AccessToken,
			token.RefreshToken,
			token.AppId,
			token.AppName,
			token.AccountId,
		})
	}

	c.JSON(http.StatusOK, res)
}