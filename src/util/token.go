package util

import (
	"strings"
	"time"

	"github.com/authink/ink.go/src/ext"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(key string, duration time.Duration, appId uint32, appName string, accountId uint32, email, uuid string) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		ext.JwtClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer: "ink.go-server",
				Audience: jwt.ClaimStrings{
					appName,
				},
				Subject:   email,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration * time.Second)),
				ID:        uuid,
			},
			AppId:     int(appId),
			AccountId: int(accountId),
		},
	)

	return t.SignedString([]byte(key))
}

func VerifyToken(key string, accessToken string) (jwtClaims *ext.JwtClaims, err error) {
	jwtClaims = &ext.JwtClaims{}

	_, err = jwt.ParseWithClaims(
		accessToken,
		jwtClaims,
		func(token *jwt.Token) (any, error) {
			return []byte(key), nil
		},
	)
	return
}

func GenerateUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
