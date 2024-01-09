package util

import (
	"crypto/sha256"
	"encoding/base64"
)

func Sha256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
