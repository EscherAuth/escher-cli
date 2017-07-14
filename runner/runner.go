package runner

import (
	"io"
	"os/exec"

	"github.com/EscherAuth/escher-cli/environment"
)

type Runner interface {
	Start(stdout, stderr io.Writer) error
	Wait() error
	EnvDiff() environment.EnvDiff
}

type runner struct {
	command *exec.Cmd
	env     environment.Environment
	diff    environment.EnvDiff
}

func New(command *exec.Cmd, env environment.Environment) Runner {
	return &runner{command: command, env: env}
}

func (r runner) EnvDiff() environment.EnvDiff {
	return r.diff
}

// Handle signals
// go func() { r.command.Process.Signal(<-sp.signal) }()

func (r *runner) Start(stdout, stderr io.Writer) error {
	var err error

	err = r.setEnvForCommand()

	if err != nil {
		return err
	}

	err = r.setPortInCommandArgs()

	if err != nil {
		return err
	}

	err = r.setOutputs(stdout, stderr)

	if err != nil {
		return err
	}

	err = r.command.Start()

	if err != nil {
		return err
	}

	return nil
}

func (r *runner) Wait() error {
	return r.command.Wait()
}

// func (sp *subProcess) envForSubProcess() []string {
// 	return sp.addNewPort(sp.removeOldPortFromEnv())
// }

// func (sp *subProcess) addNewPort(env []string) []string {
// 	env = append(env, "PORT="+sp.port)
// 	return env
// }

// func (sp *subProcess) removeOldPortFromEnv() []string {
// 	var newEnv []string
// 	for _, v := range os.Environ() {
// 		if match, _ := regexp.MatchString("PORT", v); match == false {
// 			newEnv = append(newEnv, v)
// 		}
// 	}
// 	return newEnv
// }
