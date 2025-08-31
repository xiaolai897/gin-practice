package utils

import (
	"crypto/sha256"
	"fmt"
)

func StringSha256(str string) string {
	crypt := sha256.Sum256([]byte(str))
	strCrypt := fmt.Sprintf("%x", crypt)
	return strCrypt
}
