package environment_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/EscherAuth/escher-cli/environment"
)

func TestForwardProxyAddr_LocalhostAddressReturned(t *testing.T) {
	t.Parallel()

	env := environment.New()

	resultAddr := env.ForwardProxyAddr()
	require.True(t, strings.HasPrefix(resultAddr, "localhost:"))

}

func TestForwardProxyAddr_EnvDiffWasEmpty_HTTPProxyValuesPopulated(t *testing.T) {
	t.Parallel()

	env := environment.New()

	resultAddr := env.ForwardProxyAddr()
	diff := env.EnvDifferencesForSubProcess()

	expected_keys := []string{
		"http_proxy",
		"HTTP_PROXY",
		"https_proxy",
		"HTTPS_PROXY",
	}

	for _, key := range expected_keys {
		require.Equal(t, resultAddr, diff[key])
	}

}

func TestForwardProxyAddr_NewOpenPortSetToTheAddress(t *testing.T) {
	t.Parallel()

	env := environment.New()

	resultAddr := env.ForwardProxyAddr()
	port, convErr := strconv.Atoi(strings.TrimPrefix(resultAddr, "localhost:"))

	require.Nil(t, convErr)
	require.True(t, port != 0)
}

func TestForwardProxyAddr_WasCalledAlready_ReturnTheSameValue(t *testing.T) {
	t.Parallel()

	env := environment.New()

	firstRunResultAddr := env.ForwardProxyAddr()
	firstENVValue := env.EnvDifferencesForSubProcess()["HTTP_PROXY"]

	secondRunResultAddr := env.ForwardProxyAddr()
	secondENVValue := env.EnvDifferencesForSubProcess()["HTTP_PROXY"]

	require.Equal(t, firstRunResultAddr, secondRunResultAddr)
	require.Equal(t, firstENVValue, secondENVValue)

}
