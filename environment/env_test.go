package environment_test

import (
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestEnvForChildCommandNoReplaceRequested(t *testing.T) {
	defer SetEnvForTheTest(t, "HOST", "FOO")()

	envKeyValuePairs := environment.New().EnvForChildCommand(map[string]string{})

	var found bool
	for _, keyValuePair := range envKeyValuePairs {
		if keyValuePair == "HOST=FOO" {
			found = true
		}
	}

	if !found {
		t.Fatal("env not set for the given replacement")
	}

}

func TestEnvForChildCommandChangePresentInTheCurrentEnv(t *testing.T) {
	defer SetEnvForTheTest(t, "HOST", "FOO")()

	envKeyValuePairs := environment.New().EnvForChildCommand(map[string]string{"HOST": "BAZ"})

	var found bool
	for _, keyValuePair := range envKeyValuePairs {
		if keyValuePair == "HOST=BAZ" {
			found = true
		}
	}

	if !found {
		t.Fatal("env not set for the given replacement")
	}

}

func TestEnvForChildCommandChangeMissingFromTheCurrentEnv(t *testing.T) {

	envKeyValuePairs := environment.New().EnvForChildCommand(map[string]string{"HOST": "BAZ"})

	var found bool
	for _, keyValuePair := range envKeyValuePairs {
		if keyValuePair == "HOST=BAZ" {
			found = true
		}
	}

	if !found {
		t.Fatal("env not set for the given replacement")
	}

}

func TestEnvForChildCommandOnEnvInstanceIfProxyGivenEnvWillBeSet(t *testing.T) {

	env := environment.New()

	changes := env.EnvDifferencesForSubProcess()

	expectedEnvKeyPair := "PORT=" + changes["PORT"]

	var found bool
	for _, keyValuePair := range env.EnvForChildCommand(changes) {
		if keyValuePair == expectedEnvKeyPair {
			found = true
		}
	}

	if !found {
		t.Fatal("env not set for the given replacement")
	}

}

func TestEnvForChildCommandIsCached(t *testing.T) {

	env := environment.New()

	changes1 := env.EnvDifferencesForSubProcess()
	changes2 := env.EnvDifferencesForSubProcess()

	if changes1["PORT"] != changes2["PORT"] {
		t.Fatal("changes should be matching because caching")
	}

}
