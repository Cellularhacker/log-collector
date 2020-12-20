package uCrypt

import (
	"crypto/sha1"
	"encoding/base64"
)

func GetSHA1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
