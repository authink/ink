package util

import (
	"crypto/rand"
	"encoding/base64"
)

func RandString(length int) string {
	randomBytes := make([]byte, length)

	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(randomBytes)
}
