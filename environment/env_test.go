package environment_test

import (
	"os"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
)

func TestNewProxyCreatedByEnvVariables(t *testing.T) {
	defer setEnvForTest(t, "HTTP_PROXY", "FOO")()
	defer setEnvForTest(t, "HTTPS_PROXY", "BAZ")()

	env := environment.New()

	if env.Proxy.HTTP != "FOO" {
		t.Error("HTTP_PROXY flag not parsed!")
	}

	if env.Proxy.HTTPS != "BAZ" {
		t.Error("https proxy env value not found!")
	}

}

func TestNewProxyLowerCasedEnvVariables(t *testing.T) {
	defer setEnvForTest(t, "http_proxy", "FOO")()
	defer setEnvForTest(t, "https_proxy", "BAZ")()

	env := environment.New()

	if env.Proxy.HTTP != "FOO" {
		t.Error("HTTP_PROXY flag not parsed!")
	}

	if env.Proxy.HTTPS != "BAZ" {
		t.Error("https proxy env value not found!")
	}

}

func TestNewProxyMissingHTTPSProxy(t *testing.T) {

	defer setEnvForTest(t, "http_proxy", "FOO")()
	defer setEnvForTest(t, "https_proxy", "")()

	env := environment.New()

	if env.Proxy.HTTP != "FOO" {
		t.Error("HTTP_PROXY flag not parsed!")
	}

	if env.Proxy.HTTPS != "FOO" {
		t.Error("https proxy env value not found!")
	}

}

func setEnvForTest(t testing.TB, key, value string) func() {

	orgEnvValue := os.Getenv(key)
	err := os.Setenv(key, value)

	if err != nil {
		t.Fatal(err)
	}

	return func() {
		err := os.Setenv(key, orgEnvValue)

		if err != nil {
			t.Fatal(err)
		}
	}

}
