package helpers

import (
	"crypto/sha256"
	"fmt"
)

/**
 * GetStringLength(string) int
 *
 * Returns the length of the string.
 */
func GetStringLength(str string) int {
	return len(str)
}

/**
 * CreateSHA256Hash(string) string
 *
 * Returns a SHA256 string.
 */
func CreateSHA256Hash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
