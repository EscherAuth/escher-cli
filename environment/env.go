package environment

import (
	"fmt"
	"os"
	"strings"
)

func EnvForChildCommand(replaces map[string]string) []string {
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

func (e Environment) EnvForChildCommand() ([]string, error) {
	replaces := map[string]string{}

	port, err := e.Port.FindOpenAsString()

	if err != nil {
		return nil, err
	}

	replaces["PORT"] = port

	return EnvForChildCommand(replaces), nil
}
