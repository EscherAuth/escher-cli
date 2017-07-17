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

func NewSubject(commandString ...string) (runner.Runner, *environment.Environment) {
	env := environment.New()
	cmd := exec.Command(commandString[0], commandString[1:]...)

	return runner.New(env, cmd), env
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

func TestRunnerSetPortForTheNewEnv(t *testing.T) {

	runner, env := NewSubject("env")
	stdout, _ := runAndWait(t, runner)

	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=" + env.EnvDifferencesForSubProcess()["PORT"]))

	if !rgx.Match(stdout) {
		t.Fatal("execution failed with modified env")
	}

}

func TestRunnerAccessChangesForTheCurrentlyRunningProcess(t *testing.T) {

	runner, env := NewSubject("echo", "hy")

	runAndWait(t, runner)

	if _, ok := env.EnvDifferencesForSubProcess()["PORT"]; !ok {
		t.Fatal("expected port to be set in the env")
	}

}

func TestRunnerWhenParamsIncludeAPORTThatIsTheSameAsInTheCurrentEnvSourcePortThanItWillBeReplaced(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "1234")()

	testCases := map[string]string{
		"1234":    "%v\n",
		"-p=1234": "-p=%v\n",

		"-p=12345": "-p=12345\n",
		"12345":    "12345\n",
		"01234":    "01234\n",
	}

	for cmdParameter, format := range testCases {

		runner, env := NewSubject("echo", cmdParameter)
		out, _ := runAndWait(t, runner)

		formatted := fmt.Sprintf(format, env.EnvDifferencesForSubProcess()["PORT"])
		parts := strings.SplitN(formatted, "%!", 2)

		expected := parts[0]
		actually := string(out)

		t.Logf("expected: %q, actually: %q", expected, actually)
		assert.Equal(t, expected, actually)

	}
}
