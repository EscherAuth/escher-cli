package environment_test

import (
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestNewProxyCreatedByEnvVariables(t *testing.T) {
	defer SetEnvForTheTest(t, "HTTP_PROXY", "FOO")()
	defer SetEnvForTheTest(t, "HTTPS_PROXY", "BAZ")()

	env := environment.New()

	if env.Proxy.HTTP != "FOO" {
		t.Error("HTTP_PROXY env value not found!")
	}

	if env.Proxy.HTTPS != "BAZ" {
		t.Error("https proxy env value not found!")
	}

}

func TestNewProxyLowerCasedEnvVariables(t *testing.T) {
	defer SetEnvForTheTest(t, "http_proxy", "FOO")()
	defer SetEnvForTheTest(t, "https_proxy", "BAZ")()

	env := environment.New()

	if env.Proxy.HTTP != "FOO" {
		t.Error("HTTP_PROXY env value not found!")
	}

	if env.Proxy.HTTPS != "BAZ" {
		t.Error("https proxy env value not found!")
	}

}

func TestNewProxyMissingHTTPSProxy(t *testing.T) {

	defer SetEnvForTheTest(t, "http_proxy", "FOO")()
	defer SetEnvForTheTest(t, "https_proxy", "")()

	env := environment.New()

	if env.Proxy.HTTP != "FOO" {
		t.Error("HTTP_PROXY env value not found!")
	}

	if env.Proxy.HTTPS != "FOO" {
		t.Error("https proxy env value not found!")
	}

}

func TestNewHostIsReturned(t *testing.T) {
	defer SetEnvForTheTest(t, "HOST", "FOO")()

	env := environment.New()

	if env.Host != "FOO" {
		t.Error("HOST env value not found!")
	}

}

func TestEnvForChildCommandNoReplaceRequested(t *testing.T) {
	defer SetEnvForTheTest(t, "HOST", "FOO")()

	envKeyValuePairs := environment.EnvForChildCommand(map[string]string{})

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

	envKeyValuePairs := environment.EnvForChildCommand(map[string]string{"HOST": "BAZ"})

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

	envKeyValuePairs := environment.EnvForChildCommand(map[string]string{"HOST": "BAZ"})

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
