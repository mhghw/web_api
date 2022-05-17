package utils

import (
	"crypto/sha1"
	"fmt"
)

func HashPassword(password string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}
