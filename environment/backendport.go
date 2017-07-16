package environment

import "strconv"

func (env *Environment) BackendPort() int {
	backendPort, _ := strconv.Atoi(env.EnvDifferencesForSubProcess()["PORT"])
	return backendPort
}
