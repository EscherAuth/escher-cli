package runner

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestRunnerSetAnOpenPortForTheCommandIfPortIsAlreadyDefinedForTheCurrentProcess(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "1234")()

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
