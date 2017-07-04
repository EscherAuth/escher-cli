package runner

import (
	"io"
	"os/exec"
)

type Runner interface {
	Start(stdout, stderr io.Writer) error
	Wait() error
}

type runner struct {
	command *exec.Cmd
	env     map[string]string
}

func New(command *exec.Cmd, childProcessEnv map[string]string) Runner {
	return &runner{command: command, env: childProcessEnv}
}

// Handle signals
// go func() { r.command.Process.Signal(<-sp.signal) }()

func (r *runner) Start(stdout, stderr io.Writer) error {

	r.setEnvForCommand()

	err := r.setOutputs(stdout, stderr)

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
