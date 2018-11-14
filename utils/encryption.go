package utils

import (
	"auxpi/bootstrap"
	"crypto/sha256"
	"fmt"
)

func GetSha256CodeWithSalt(s string) string {
	h := sha256.New()
	salt := bootstrap.SiteConfig.AuxpiSalt
	h.Write([]byte(s+salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}
