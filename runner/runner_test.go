package runner

import (
	"bytes"
	"os/exec"
	"regexp"
	"testing"

	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestCommand(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "1234")()

	cmd := exec.Command("env")
	env := map[string]string{"PORT": "4321"}
	r := New(cmd, env)

	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})

	err := r.Start(stdout, stderr)

	if err != nil {
		t.Fatal(err)
	}

	err = r.Wait()

	if err != nil {
		t.Fatal(err)
	}

	rgx := regexp.MustCompile(regexp.QuoteMeta("PORT=4321"))

	if !rgx.Match(stdout.Bytes()) {
		t.Fatal("execution failed with modified env")
	}

}
