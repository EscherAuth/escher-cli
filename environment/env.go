package environment

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvDiff map[string]string

func (e Environment) EnvForChildCommand(replaces EnvDiff) []string {
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

func (e Environment) EnvDifferencesForSubProcess() (EnvDiff, error) {
	changes := make(EnvDiff)

	sourcePort, sourcePortIsGiven := e.Port.Source()

	var port int

	for {
		port = RequestPortFromOperationSystem()

		if !sourcePortIsGiven {
			break
		}

		if port != sourcePort {
			break
		}
	}

	changes["PORT"] = strconv.Itoa(port)

	return changes, nil

}
