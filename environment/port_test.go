package environment_test

import (
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestSourcePortSetAndFound(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "3000")()

	env := environment.New()

	port, found := env.Port.Source()

	if !found {
		t.Fatal("expected port value not found")
	}

	if port != "3000" {
		t.Fatal("not expected port value: " + port)
	}

}

func TestSourcePortNotSetInTheEnv(t *testing.T) {

	env := environment.New()

	_, found := env.Port.Source()

	if found {
		t.Fatal("expected port value to be missing")
	}

}

func TestFindOpenAsString(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "3000")()

	env := environment.New()

	port, notFoundErr := env.Port.FindOpenAsString()

	if notFoundErr != nil {
		t.Fatal(notFoundErr)
	}

	if port != "3000" {
		t.Fatal("not expected port value: " + port)
	}

}
