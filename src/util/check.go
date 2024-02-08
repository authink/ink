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

func CheckApp(extCtx *ext.Context, err error, active bool, checkSecret CheckSecretFunc, code int) (ok bool) {
	if err != nil || !active || !checkSecret() {
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			extCtx.AbortWithServerError(err)
			return
		}

		switch code {
		case http.StatusUnauthorized:
			extCtx.AbortWithUnauthorized(ext.ERR_INVALID_APP)
		default:
			extCtx.AbortWithClientError(ext.ERR_INVALID_APP)
		}

		return
	}
	return true
}

type CheckPasswordFunc func() bool

func CheckStaff(extCtx *ext.Context, err error, active, departure bool, checkPassword CheckPasswordFunc, code int) (ok bool) {
	if err != nil || !active || departure || !checkPassword() {
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			extCtx.AbortWithServerError(err)
			return
		}

		switch code {
		case http.StatusUnauthorized:
			extCtx.AbortWithUnauthorized(ext.ERR_INVALID_ACCOUNT)
		default:
			extCtx.AbortWithClientError(ext.ERR_INVALID_ACCOUNT)
		}

		return
	}
	return true
}

func CheckAccessToken(extCtx *ext.Context, secretKey, accessToken, uuid string) (jwtClaims *ext.JwtClaims, ok bool) {
	jwtClaims, err := VerifyToken(secretKey, accessToken)

	if (err != nil && !errors.Is(err, jwt.ErrTokenExpired)) || jwtClaims.ID != uuid {
		extCtx.AbortWithClientError(ext.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	ok = true
	return
}
