package environment

import "strconv"

type EnvDiff map[string]string

func (e *Environment) EnvDifferencesForSubProcess() EnvDiff {
	if len(e.envDifferencesForSubProcess) == 0 {
		e.envDifferencesForSubProcess = e.createEnvDifferencesForSubProcess()
	}

	return e.envDifferencesForSubProcess
}

func (e *Environment) createEnvDifferencesForSubProcess() EnvDiff {
	changes := make(EnvDiff)
	e.setNewUsablePortForEnvDiff(changes)
	// e.setNewUsableProxyPortForEnvDiff()
	return changes
}

func (e *Environment) setNewUsablePortForEnvDiff(changes EnvDiff) {
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
}
