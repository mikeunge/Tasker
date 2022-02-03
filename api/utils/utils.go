package utils

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("key not set")
	} else {
		return val, nil
	}
}
