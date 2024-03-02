package util

import (
	"database/sql"
	errs "errors"
	"net/http"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/inkstone/util"
	"github.com/authink/inkstone/web"
	"github.com/golang-jwt/jwt/v5"
)

type CheckSecretFunc func() bool

func CompareSecrets(secret, reqAppSecret string) (ok bool) {
	return secret == util.Sha256(reqAppSecret)
}

func CheckApp(c *web.Context, err error, active bool, checkSecret CheckSecretFunc, code int) (ok bool) {
	if err != nil || !active || !checkSecret() {
		if err != nil && !errs.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}

		switch code {
		case http.StatusUnauthorized:
			c.AbortWithUnauthorized(errors.ERR_INVALID_APP)
		default:
			c.AbortWithClientError(errors.ERR_INVALID_APP)
		}

		return
	}
	return true
}

type CheckPasswordFunc func() bool

func CheckStaff(c *web.Context, err error, active, departure bool, checkPassword CheckPasswordFunc, code int) (ok bool) {
	if err != nil || !active || departure || !checkPassword() {
		if err != nil && !errs.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}

		switch code {
		case http.StatusUnauthorized:
			c.AbortWithUnauthorized(errors.ERR_INVALID_ACCOUNT)
		default:
			c.AbortWithClientError(errors.ERR_INVALID_ACCOUNT)
		}

		return
	}
	return true
}

func CheckAccessToken(c *web.Context, secretKey, accessToken, uuid string) (jwtClaims *util.JwtClaims, ok bool) {
	jwtClaims, err := util.VerifyToken(secretKey, accessToken)

	if (err != nil && !errs.Is(err, jwt.ErrTokenExpired)) || jwtClaims.ID != uuid {
		c.AbortWithClientError(errors.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	ok = true
	return
}
