package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthN(c *gin.Context) {
	extCtx := (*ext.Context)(c)
	ink := c.MustGet("ink").(*core.Ink)

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		extCtx.AbortWithUnauthorized(ext.ERR_MISSING_ACCESS_TOKEN)
		return
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := util.VerifyToken(ink.Env.SecretKey, accessToken)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			extCtx.AbortWithUnauthorized(ext.ERR_EXPIRED_ACCESS_TOKEN)
			return
		}

		extCtx.AbortWithUnauthorized(ext.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	if _, err = ink.GetByAccessToken(claims.ID); err != nil {
		extCtx.AbortWithUnauthorized(ext.ERR_REVOKED_ACCESS_TOKEN)
		return
	}

	app, err := ink.GetApp(claims.AppId)
	if !util.CheckApp(extCtx, err, app.Active, func() bool { return true }, http.StatusUnauthorized) {
		return
	}
	c.Set("app", app)

	switch app.Name {
	case core.APP_ADMIN_DEV:
		staff, err := ink.GetStaff(claims.AccountId)

		if ok := util.CheckStaff(extCtx, err, staff.Active, staff.Departure, func() bool { return true }, http.StatusUnauthorized); !ok {
			return
		}

		staff.Password = ""
		c.Set("account", staff)

	default:
		extCtx.AbortWithUnauthorized(ext.ERR_UNSUPPORTED_APP)
		return
	}

	c.Next()
}
