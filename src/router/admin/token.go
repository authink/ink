package admin

import (
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type pageReq struct {
	Offset int `form:"offset" binding:"min=0"`
	Limit  int `form:"limit" binding:"required,min=1,max=100"`
}

type pageRes[T any] struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Items  []T `json:"items,omitempty"`
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
//	@Param			offset	query		int	false	"offset"
//	@Param			limit	query		int	true	"limit"
//	@Success		200		{object}	pageRes[tokenRes]
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

	var total int
	var tokens []model.AuthTokenWithApp

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		if total, err = orm.AuthToken(appContext).CountWithTx(tx); err != nil {
			return
		}

		tokens, err = orm.AuthToken(appContext).PaginationWithTx(req.Offset, req.Limit, tx)
		return
	}); err != nil {
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

	c.Response(&pageRes[tokenRes]{
		total,
		req.Offset,
		req.Limit,
		res,
	})
}

type delTokenReq struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

// deleteToken godoc
//
//	@Summary		Delete a token
//	@Description	Delete a token
//	@Tags			admin_token
//	@Router			/admin/tokens/{id}	[delete]
//	@Security		ApiKeyAuth
//	@Param			id	path		int	true	"token id"
//	@Success		200	{string}	empty
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func deleteToken(c *inkstone.Context) {
	req := &delTokenReq{}
	if err := c.ShouldBindUri(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	if err := orm.AuthToken(c.App()).Delete(req.Id); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
