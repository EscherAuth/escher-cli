package environment_test

import (
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestHTTPProxyValueSetAndFound(t *testing.T) {
	defer SetEnvForTheTest(t, "HTTP_PROXY", "FOO")()

	env := environment.New()

	http, found := env.Proxy.HTTP()

	if !found {
		t.Fatal("http proxy en should be found")
	}

	if http != "FOO" {
		t.Fatal("HTTP_PROXY env value not found!")
	}

}

func TestHTTPSProxyValueSetAndFound(t *testing.T) {
	defer SetEnvForTheTest(t, "HTTPS_PROXY", "BAZ")()

	env := environment.New()

	https, found := env.Proxy.HTTPS()

	if !found {
		t.Fatal("http proxy en should be found")
	}

	if https != "BAZ" {
		t.Fatal("HTTP_PROXY env value not found!")
	}

}

func TestProxyKeysAreInLowerCasedEnvVariables(t *testing.T) {
	defer SetEnvForTheTest(t, "http_proxy", "FOO")()
	defer SetEnvForTheTest(t, "https_proxy", "BAZ")()

	env := environment.New()
	http, _ := env.Proxy.HTTP()
	https, _ := env.Proxy.HTTPS()

	if http != "FOO" {
		t.Fatal("HTTP_PROXY env value not found!")
	}

	if https != "BAZ" {
		t.Fatal("https proxy env value not found!")
	}

}

func TestProxyHTTPSValueIsMissingAndHTTPValueUsed(t *testing.T) {
	defer SetEnvForTheTest(t, "http_proxy", "FOO")()

	env := environment.New()
	https, found := env.Proxy.HTTPS()

	if !found {
		t.Fatal("expected to use http_proxy env value")
	}

	if https != "FOO" {
		t.Fatal("HTTP_PROXY env value not found!")
	}

}
