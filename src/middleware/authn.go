package middleware

import (
	errs "errors"
	"net/http"
	"strings"

	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone"
	"github.com/golang-jwt/jwt/v5"
)

func AuthN(c *inkstone.Context) {
	appCtx := c.AppContext()
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithUnauthorized(errors.ERR_MISSING_ACCESS_TOKEN)
		return
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := inkstone.VerifyToken(appCtx.SecretKey, accessToken)
	if err != nil {
		if errs.Is(err, jwt.ErrTokenExpired) {
			c.AbortWithUnauthorized(errors.ERR_EXPIRED_ACCESS_TOKEN)
			return
		}

		c.AbortWithUnauthorized(errors.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	if _, err = orm.AuthToken(appCtx).GetByAccessToken(claims.ID); err != nil {
		c.AbortWithUnauthorized(errors.ERR_REVOKED_ACCESS_TOKEN)
		return
	}

	app, err := orm.App(appCtx).Get(claims.AppId)
	if !util.CheckApp(c, err, app.Active, func() bool { return true }, http.StatusUnauthorized) {
		return
	}
	c.Set("app", app)

	switch app.Name {
	case env.AppNameAdmin():
		staff, err := orm.Staff(appCtx).Get(claims.AccountId)

		if ok := util.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return true }, http.StatusUnauthorized); !ok {
			return
		}

		staff.Password = ""
		c.Set("account", staff)

	default:
		c.AbortWithUnauthorized(errors.ERR_UNSUPPORTED_APP)
		return
	}

	c.Next()
}
