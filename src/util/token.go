package util

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(key []byte, appId uint32, appName string, accountId uint32, email, uuid string) (string, error) {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "ink.go-server",
			"aud": appName,
			"sub": email,
			"iat": time.Now().Unix(),
			// todo move to Env
			"exp":       time.Now().Add(2 * time.Hour).Unix(),
			"appId":     appId,
			"accountId": accountId,
			"uuid":      uuid,
		},
	)

	return t.SignedString(key)
}

func GenerateUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
