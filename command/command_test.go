package command_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/EscherAuth/escher-cli/command"
	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
	"github.com/stretchr/testify/assert"
)

func NewSubject(t testing.TB, commandString ...string) (*environment.Environment, *exec.Cmd) {
	env := environment.New()
	cmd := exec.Command(commandString[0], commandString[1:]...)
	err := command.ConfigureBy(env, cmd)

	if err != nil {
		t.Fatal(err)
	}

	return env, cmd
}

func runAndWait(t testing.TB, cmd *exec.Cmd) ([]byte, []byte) {

	cmd.Stdout = bytes.NewBuffer([]byte{})
	cmd.Stderr = bytes.NewBuffer([]byte{})

	err := cmd.Run()

	if err != nil {
		t.Fatal(err)
	}

	stdout := cmd.Stdout.(*bytes.Buffer)
	stderr := cmd.Stdout.(*bytes.Buffer)

	return stdout.Bytes(), stderr.Bytes()

}

func TestRunnerSetPortForTheNewEnv(t *testing.T) {

	env, cmd := NewSubject(t, "env")
	stdout, _ := runAndWait(t, cmd)

	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=" + env.EnvDifferencesForSubProcess()["PORT"]))

	if !rgx.Match(stdout) {
		t.Fatalf("execution failed with modified env\n%s", stdout)
	}

}

func TestRunnerAccessChangesForTheCurrentlyRunningProcess(t *testing.T) {

	env, cmd := NewSubject(t, "echo", "hy")

	runAndWait(t, cmd)

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

		env, cmd := NewSubject(t, "echo", cmdParameter)
		out, _ := runAndWait(t, cmd)

		formatted := fmt.Sprintf(format, env.EnvDifferencesForSubProcess()["PORT"])
		parts := strings.SplitN(formatted, "%!", 2)

		expected := parts[0]
		actually := string(out)

		t.Logf("expected: %q, actually: %q", expected, actually)
		assert.Equal(t, expected, actually)

	}
}
