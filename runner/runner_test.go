package runner_test

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/EscherAuth/escher-cli/runner"
)

func TestRunnerSetPortForTheNewEnv(t *testing.T) {

	cmd := exec.Command("env")
	env := environment.New()

	r := runner.New(cmd, env)
	stdout, _ := runAndWait(t, r)

	portAsString, _ := env.Port.FindOpenAsString()
	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=" + portAsString))

	if !rgx.Match(stdout) {
		t.Fatal("execution failed with modified env")
	}

}

func TestRunnerAccessChangesForTheCcurrentlyRunningProcess(t *testing.T) {

	cmd := exec.Command("echo", "hy")
	env := environment.New()
	r := runner.New(cmd, env)
	runAndWait(t, r)

	diff, _ := env.EnvDifferencesForSubProcess()
	processEnvDiff := r.EnvDiff()

	portGivenToProcess, ok := processEnvDiff["PORT"]

	if !ok {
		t.Fatal("Missing Port value about the changes")
	}

	if portGivenToProcess != diff["PORT"] {
		t.Fatal("expected port is different from the found one")
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
