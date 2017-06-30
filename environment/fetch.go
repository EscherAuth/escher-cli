package environment

import (
	"os"
	"strings"
)

func fetchEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		value = os.Getenv(strings.ToUpper(key))
	}

	if value == "" {
		value = os.Getenv(strings.ToLower(key))
	}

	return value
}
