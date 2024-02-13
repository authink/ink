package util

import (
	"errors"
	"net/http"

	"database/sql"

	"github.com/authink/ink.go/src/ext"
	"github.com/golang-jwt/jwt/v5"
)

type CheckSecretFunc func() bool

func CompareSecrets(secret, reqAppSecret string) (ok bool) {
	return secret == Sha256(reqAppSecret)
}

func CheckApp(c *ext.Context, err error, active bool, checkSecret CheckSecretFunc, code int) (ok bool) {
	if err != nil || !active || !checkSecret() {
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}

		switch code {
		case http.StatusUnauthorized:
			c.AbortWithUnauthorized(ext.ERR_INVALID_APP)
		default:
			c.AbortWithClientError(ext.ERR_INVALID_APP)
		}

		return
	}
	return true
}

type CheckPasswordFunc func() bool

func CheckStaff(c *ext.Context, err error, active, departure bool, checkPassword CheckPasswordFunc, code int) (ok bool) {
	if err != nil || !active || departure || !checkPassword() {
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}

		switch code {
		case http.StatusUnauthorized:
			c.AbortWithUnauthorized(ext.ERR_INVALID_ACCOUNT)
		default:
			c.AbortWithClientError(ext.ERR_INVALID_ACCOUNT)
		}

		return
	}
	return true
}

func CheckAccessToken(c *ext.Context, secretKey, accessToken, uuid string) (jwtClaims *ext.JwtClaims, ok bool) {
	jwtClaims, err := VerifyToken(secretKey, accessToken)

	if (err != nil && !errors.Is(err, jwt.ErrTokenExpired)) || jwtClaims.ID != uuid {
		c.AbortWithClientError(ext.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	ok = true
	return
}
