package token

import (
	libsql "database/sql"
	"errors"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	//"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
)

type reqRefresh struct {
	AccessToken  string `json:"access_token" binding:"required,min=1"`
	RefreshToken string `json:"refresh_token" binding:"required,min=1"`
}

func refresh(c *gin.Context) {
	req := &reqRefresh{}
	if err := c.BindJSON(req); err != nil {
		return
	}

	extCtx := (*ext.Context)(c)
	ink := c.MustGet("ink").(*core.Ink)

	var authToken model.AuthToken
	if err := ink.DB.Get(
		&authToken,
		sql.Query.GetAuthTokenByRefreshToken,
		req.RefreshToken,
	); err != nil {
		if !errors.Is(err, libsql.ErrNoRows) {
			extCtx.AbortWithServerError(err)
			return
		}
		extCtx.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	// jwtClaims, err := util.VerifyToken(
	// 	[]byte(ink.Env.SecretKey),
	// 	authToken.AccessToken,
	// )
}
