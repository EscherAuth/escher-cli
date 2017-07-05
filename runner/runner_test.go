package runner

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
)

func TestRunnerSetPortForTheNewEnv(t *testing.T) {

	cmd := exec.Command("env")
	env := environment.New()

	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})

	r := New(cmd, env)
	err := r.Start(stdout, stderr)

	if err != nil {
		t.Fatal(err)
	}

	err = r.Wait()

	if err != nil {
		t.Fatal(err)
	}

	portAsString, _ := env.Port.FindOpenAsString()
	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=" + portAsString))

	if !rgx.Match(stdout.Bytes()) {
		t.Fatal("execution failed with modified env")
	}

}
