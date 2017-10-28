package environment

import (
	"fmt"
	"os"
	"strings"
)

func (e *Environment) EnvForChildCommand(replaces EnvDiff) []string {

	var env []string

	for _, keyValuePair := range os.Environ() {

		key := strings.Split(keyValuePair, "=")[0]

		if _, ok := replaces[key]; !ok {
			env = append(env, keyValuePair)
		}

	}

	for k, v := range replaces {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	return env
}
