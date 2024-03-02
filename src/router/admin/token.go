package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	o "github.com/authink/inkstone/orm"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func setupTokenGroup(gAdmin *gin.RouterGroup) {
	gTokens := gAdmin.Group(authz.Tokens.Name)
	gTokens.Use(middleware.Authz(authz.Tokens))
	gTokens.GET("", web.HandlerAdapter(tokens))
	gTokens.DELETE(":id", web.HandlerAdapter(deleteToken))
}

type tokenRes struct {
	web.Response
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
//	@Success		200		{object}	web.PagingResponse[tokenRes]
//	@Failure		400		{object}	web.ClientError
//	@Failure		401		{object}	web.ClientError
//	@Failure		403		{object}	web.ClientError
//	@Failure		500		{string}	empty
func tokens(c *web.Context) {
	appCtx := c.AppContext()

	req := new(web.PagingRequest)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var total int
	var tokens []model.AuthTokenWithApp

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		if total, err = orm.AuthToken(appCtx).CountTx(tx); err != nil {
			return
		}

		pageArgs := o.PageArgs{
			Offset: req.Offset,
			Limit:  req.Limit,
		}

		tokens, err = orm.AuthToken(appCtx).PaginationTx(tx, &pageArgs)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res = []tokenRes{}
	for i := range tokens {
		token := &tokens[i]
		res = append(res, tokenRes{
			web.Response{
				Id:        int(token.Id),
				CreatedAt: token.CreatedAt,
			},
			token.AccessToken,
			token.RefreshToken,
			int(token.AppId),
			token.AppName,
			int(token.AccountId),
		})
	}

	c.Response(&web.PagingResponse[tokenRes]{
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
//	@Failure		400	{object}	web.ClientError
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func deleteToken(c *web.Context) {
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
