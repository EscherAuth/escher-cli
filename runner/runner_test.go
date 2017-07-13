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

	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=" + r.EnvDiff()["PORT"]))

	if !rgx.Match(stdout) {
		t.Fatal("execution failed with modified env")
	}

}

func TestRunnerAccessChangesForTheCurrentlyRunningProcess(t *testing.T) {

	cmd := exec.Command("echo", "hy")
	env := environment.New()
	r := runner.New(cmd, env)
	runAndWait(t, r)

	envDiff := r.EnvDiff()

	if _, ok := envDiff["PORT"]; !ok {
		t.Fatal("expected port to be set in the env")
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
