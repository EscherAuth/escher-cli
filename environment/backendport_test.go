package environment_test

import (
	"strconv"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/stretchr/testify/assert"
)

func TestBackendPort(t *testing.T) {

	env := environment.New()

	diff := env.EnvDifferencesForSubProcess()

	expected, err := strconv.Atoi(diff["PORT"])

	if err != nil {
		t.Fatal(err)
	}

	actually := env.BackendPort()

	assert.Equal(t, expected, actually, "backend port is not equal!")

}
