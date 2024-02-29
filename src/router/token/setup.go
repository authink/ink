package token

import (
	"database/sql"
	errs "errors"
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func SetupTokenGroup(rg *gin.RouterGroup) {
	gToken := rg.Group("token")
	gToken.POST("grant", inkstone.HandlerAdapter(grant))
	gToken.POST("refresh", inkstone.HandlerAdapter(refresh))
	gToken.POST("revoke", inkstone.HandlerAdapter(revoke))
}

func generateAuthToken(c *inkstone.Context, app *model.App, staff *model.Staff) (res *GrantRes) {
	appCtx := c.AppContext()

	jwtClaims := inkstone.NewJwtClaims(
		inkstone.GenerateUUID(),
		appCtx.AppName,
		app.Name,
		staff.Email,
		time.Duration(appCtx.AccessTokenDuration),
		app.Id,
		staff.Id,
	)

	accessToken, err := inkstone.GenerateToken(
		appCtx.SecretKey,
		jwtClaims,
	)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	refreshToken := inkstone.GenerateUUID()

	authToken := model.NewAuthToken(jwtClaims.ID, refreshToken, app.Id, staff.Id)

	if err = orm.AuthToken(appCtx).Insert(authToken); err != nil {
		c.AbortWithServerError(err)
		return
	}

	res = &GrantRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(appCtx.AccessTokenDuration),
	}
	return
}

func checkRefreshToken(c *inkstone.Context, refreshToken string) (authToken *model.AuthToken, ok bool) {
	appCtx := c.AppContext()
	authToken, err := orm.AuthToken(appCtx).GetByRefreshToken(refreshToken)
	if err != nil {
		if !errs.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}
		c.AbortWithClientError(errors.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	if err = orm.AuthToken(appCtx).Delete(int(authToken.Id)); err != nil {
		c.AbortWithServerError(err)
		return
	}

	if time.Now().After(authToken.CreatedAt.Add(time.Duration(appCtx.RefreshTokenDuration) * time.Hour)) {
		c.AbortWithClientError(errors.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	ok = true
	return
}
