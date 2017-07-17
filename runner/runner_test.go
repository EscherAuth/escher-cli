package runner_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
	"github.com/EscherAuth/escher-cli/runner"
	"github.com/stretchr/testify/assert"
)

func TestRunnerSetPortForTheNewEnv(t *testing.T) {

	cmd := exec.Command("env")
	env := environment.New()

	r := runner.New(env, cmd)
	stdout, _ := runAndWait(t, r)

	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=" + env.EnvDifferencesForSubProcess()["PORT"]))

	if !rgx.Match(stdout) {
		t.Fatal("execution failed with modified env")
	}

}

func TestRunnerAccessChangesForTheCurrentlyRunningProcess(t *testing.T) {

	cmd := exec.Command("echo", "hy")
	env := environment.New()
	r := runner.New(env, cmd)
	runAndWait(t, r)

	if _, ok := env.EnvDifferencesForSubProcess()["PORT"]; !ok {
		t.Fatal("expected port to be set in the env")
	}

}

func TestRunnerWhenParamsIncludeAPORTThatIsTheSameAsInTheCurrentEnvSourcePortThanItWillBeReplaced(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "1234")()

	env := environment.New()

	testCases := map[string]string{
		"1234":    "%v\n",
		"-p=1234": "-p=%v\n",

		"-p=12345": "-p=12345\n",
		"12345":    "12345\n",
		"01234":    "01234\n",
	}

	for cmdParameter, format := range testCases {

		cmd := exec.Command("echo", cmdParameter)
		r := runner.New(env, cmd)
		out, _ := runAndWait(t, r)

		formatted := fmt.Sprintf(format, env.EnvDifferencesForSubProcess()["PORT"])
		parts := strings.SplitN(formatted, "%!", 2)

		expected := parts[0]
		actually := string(out)

		t.Logf("expected: %q, actually: %q", expected, actually)
		assert.Equal(t, expected, actually)

	}
}

func runAndWait(t testing.TB, r runner.Runner) (stdout []byte, stderr []byte) {

	stdoutBuffer := bytes.NewBuffer([]byte{})
	stderrBuffer := bytes.NewBuffer([]byte{})

	err := r.Start(stdoutBuffer, stderrBuffer)

	if err != nil {
		t.Fatal(err)
	}

	err = r.Wait()

	if err != nil {
		t.Fatal(err)
	}

	return stdoutBuffer.Bytes(), stderrBuffer.Bytes()

}
