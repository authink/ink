package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func setupTokenGroup(gAdmin *gin.RouterGroup) {
	gTokens := gAdmin.Group(authz.Tokens.Name)
	gTokens.Use(middleware.Authz(authz.Tokens))
	gTokens.GET("", inkstone.HandlerAdapter(tokens))
	gTokens.DELETE(":id", inkstone.HandlerAdapter(deleteToken))
}

type tokenRes struct {
	inkstone.Response
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	AppId        int    `json:"appId,omitempty"`
	AppName      string `json:"appName,omitempty"`
	AccountId    int    `json:"accountId,omitempty"`
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
//	@Success		200		{object}	inkstone.PagingResponse[tokenRes]
//	@Failure		400		{object}	inkstone.ClientError
//	@Failure		401		{object}	inkstone.ClientError
//	@Failure		403		{object}	inkstone.ClientError
//	@Failure		500		{string}	empty
func tokens(c *inkstone.Context) {
	appCtx := c.AppContext()

	req := new(inkstone.PagingRequest)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var total int
	var tokens []model.AuthTokenWithApp

	if err := inkstone.Transaction(appCtx, func(tx *sqlx.Tx) (err error) {
		if total, err = orm.AuthToken(appCtx).CountWithTx(tx); err != nil {
			return
		}

		tokens, err = orm.AuthToken(appCtx).PaginationWithTx(req.Offset, req.Limit, tx)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []tokenRes
	for i := range tokens {
		token := &tokens[i]
		res = append(res, tokenRes{
			inkstone.Response{
				Id:        int(token.Id),
				CreatedAt: token.CreatedAt,
				UpdatedAt: token.UpdatedAt,
			},
			token.AccessToken,
			token.RefreshToken,
			int(token.AppId),
			token.AppName,
			int(token.AccountId),
		})
	}

	c.Response(&inkstone.PagingResponse[tokenRes]{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  total,
		Items:  res,
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
//	@Failure		400	{object}	inkstone.ClientError
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func deleteToken(c *inkstone.Context) {
	req := new(delTokenReq)
	if err := c.ShouldBindUri(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	if err := orm.AuthToken(c.AppContext()).Delete(req.Id); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
