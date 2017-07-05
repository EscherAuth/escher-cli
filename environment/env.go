package environment

import (
	"fmt"
	"os"
	"strings"
)

type EnvChanges map[string]string

func (e Environment) EnvForChildCommand(replaces EnvChanges) []string {
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

func (e Environment) EnvDifferencesForSubProcess() (EnvChanges, error) {
	changes := make(EnvChanges)

	port, err := e.Port.FindOpenAsString()

	if err != nil {
		return nil, err
	}

	changes["PORT"] = port

	return changes, nil

}
