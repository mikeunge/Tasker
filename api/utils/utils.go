package utils

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid"
)

func GetEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("key not set")
	} else {
		return val, nil
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
