package environment

import (
	"os"
	"strings"
)

func fetchEnv(key string) (value string, found bool) {
	value, found = os.LookupEnv(key)

	if !found {
		value, found = os.LookupEnv(strings.ToUpper(key))
	}

	if !found {
		value, found = os.LookupEnv(strings.ToLower(key))
	}

	return
}
