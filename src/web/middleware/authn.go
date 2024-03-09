package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/authink/ink/src/envs"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/ink/src/web/helper"
	"github.com/authink/inkstone/jwtx"
	"github.com/authink/inkstone/web"
	"github.com/golang-jwt/jwt/v5"
)

func Authn(c *web.Context) {
	appCtx := c.AppContext()
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithUnauthorized(errs.ERR_MISSING_ACCESS_TOKEN)
		return
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := jwtx.VerifyToken(appCtx.SecretKey, accessToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.AbortWithUnauthorized(errs.ERR_EXPIRED_ACCESS_TOKEN)
			return
		}

		c.AbortWithUnauthorized(errs.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	if err = orm.AuthToken(appCtx).GetByAccessToken(&models.AuthToken{
		AccessToken: claims.ID,
	}); err != nil {
		c.AbortWithUnauthorized(errs.ERR_REVOKED_ACCESS_TOKEN)
		return
	}

	var app models.App
	app.Id = uint32(claims.AppId)

	err = orm.App(appCtx).Get(&app)
	if !helper.CheckApp(c, err, app.Active, func() bool { return true }, http.StatusUnauthorized) {
		return
	}
	c.Set("app", &app)

	switch app.Name {
	case envs.AppNameAdmin():
		var staff models.Staff
		staff.Id = uint32(claims.AccountId)

		err = orm.Staff(appCtx).Get(&staff)
		if ok := helper.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return true }, http.StatusUnauthorized); !ok {
			return
		}

		staff.Password = ""
		c.Set("account", &staff)

	default:
		c.AbortWithUnauthorized(errs.ERR_UNSUPPORTED_APP)
		return
	}

	c.Next()
}
