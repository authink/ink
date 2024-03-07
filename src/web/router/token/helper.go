package token

import (
	"database/sql"
	"errors"
	"time"

	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/web/errs"
	"github.com/authink/inkstone/jwtx"
	"github.com/authink/inkstone/util"
	"github.com/authink/inkstone/web"
)

func generateAuthToken(c *web.Context, app *models.App, staff *models.Staff) (res *GrantRes) {
	appCtx := c.AppContext()

	jwtClaims := jwtx.NewJwtClaims(
		util.GenerateUUID(),
		appCtx.AppName,
		app.Name,
		staff.Email,
		time.Duration(appCtx.AccessTokenDuration),
		app.Id,
		staff.Id,
	)

	accessToken, err := jwtx.GenerateToken(
		appCtx.SecretKey,
		jwtClaims,
	)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	refreshToken := util.GenerateUUID()

	authToken := models.NewAuthToken(jwtClaims.ID, refreshToken, app.Id, staff.Id)

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

func checkRefreshToken(c *web.Context, refreshToken string) (authToken *models.AuthToken, ok bool) {
	var appCtx = c.AppContext()
	authToken = &models.AuthToken{
		RefreshToken: refreshToken,
	}

	err := orm.AuthToken(appCtx).GetByRefreshToken(authToken)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}
		c.AbortWithClientError(errs.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	if err = orm.AuthToken(appCtx).Delete(authToken); err != nil {
		c.AbortWithServerError(err)
		return
	}

	if time.Now().After(authToken.CreatedAt.Add(time.Duration(appCtx.RefreshTokenDuration) * time.Hour)) {
		c.AbortWithClientError(errs.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	ok = true
	return
}
