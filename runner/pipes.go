package runner

import (
	"io"
)

func (r *runner) setOutputs(stdout, stderr io.Writer) error {

	cmdStdout, err := r.command.StdoutPipe()

	if err != nil {
		return err
	}

	cmdStderr, err := r.command.StderrPipe()

	if err != nil {
		return err
	}

	go io.Copy(stdout, cmdStdout)
	go io.Copy(stderr, cmdStderr)

	return nil

}
