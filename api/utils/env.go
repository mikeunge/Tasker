package utils

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf(fmt.Sprintf("%s not set\n", key))
	} else {
		return val, nil
	}
}
