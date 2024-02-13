package token

import (
	"database/sql"
	"errors"
	"time"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
)

func SetupTokenGroup(rg *gin.RouterGroup) {
	gToken := rg.Group("token")
	gToken.POST("grant", ext.Handler(grant))
	gToken.POST("refresh", ext.Handler(refresh))
	gToken.POST("revoke", ext.Handler(revoke))
}

func generateAuthToken(c *ext.Context, ink *core.Ink, app *model.App, staff *model.Staff) (res *grantRes) {
	uuid := util.GenerateUUID()
	accessToken, err := util.GenerateToken(ink.Env.SecretKey, time.Duration(ink.Env.AccessTokenDuration), app.Id, app.Name, staff.Id, staff.Email, uuid)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	refreshToken := util.GenerateUUID()
	// accessToken identified by uuid
	authToken := model.NewAuthToken(uuid, refreshToken, app.Id, staff.Id)

	if _, err = orm.AuthToken(ink).Save(authToken); err != nil {
		c.AbortWithServerError(err)
		return
	}

	res = &grantRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(ink.Env.AccessTokenDuration),
	}
	return
}

func checkRefreshToken(c *ext.Context, ink *core.Ink, refreshToken string) (authToken *model.AuthToken, ok bool) {
	authToken, err := orm.AuthToken(ink).GetByRefreshToken(refreshToken)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}
		c.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	if _, err = orm.AuthToken(ink).Delete(int(authToken.Id)); err != nil {
		c.AbortWithServerError(err)
		return
	}

	if time.Now().After(authToken.CreatedAt.Add(time.Duration(ink.Env.RefreshTokenDuration) * time.Hour)) {
		c.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	ok = true
	return
}
