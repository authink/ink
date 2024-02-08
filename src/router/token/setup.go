package token

import (
	"database/sql"
	"errors"
	"time"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
)

func SetupTokenGroup(rg *gin.RouterGroup) {
	gToken := rg.Group("token")
	gToken.POST("grant", grant)
	gToken.POST("refresh", refresh)
	gToken.POST("revoke", revoke)
}

func generateAuthToken(extCtx *ext.Context, ink *core.Ink, app *model.App, staff *model.Staff) (res *resGrant) {
	uuid := util.GenerateUUID()
	accessToken, err := util.GenerateToken(ink.Env.SecretKey, time.Duration(ink.Env.AccessTokenDuration), app.Id, app.Name, staff.Id, staff.Email, uuid)
	if err != nil {
		extCtx.AbortWithServerError(err)
		return
	}

	refreshToken := util.GenerateUUID()
	// accessToken identified by uuid
	authToken := model.NewAuthToken(uuid, refreshToken, app.Id, staff.Id)

	if _, err = ink.SaveToken(authToken); err != nil {
		extCtx.AbortWithServerError(err)
		return
	}

	res = &resGrant{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(ink.Env.AccessTokenDuration),
	}
	return
}

func checkRefreshToken(extCtx *ext.Context, ink *core.Ink, refreshToken string) (authToken *model.AuthToken, ok bool) {
	authToken, err := ink.GetByRefreshToken(refreshToken)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			extCtx.AbortWithServerError(err)
			return
		}
		extCtx.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	if _, err = ink.DeleteToken(int(authToken.Id)); err != nil {
		extCtx.AbortWithServerError(err)
		return
	}

	if time.Now().After(authToken.CreatedAt.Add(time.Duration(ink.Env.RefreshTokenDuration) * time.Hour)) {
		extCtx.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	ok = true
	return
}
