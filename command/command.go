package command

import (
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/EscherAuth/escher-cli/environment"
)

func ConfigureBy(env *environment.Environment, cmd *exec.Cmd) error {
	setCommandEnvironment(env, cmd)

	err := setCommandArgs(env, cmd)

	if err != nil {
		return err
	}

	return nil
}

func setCommandEnvironment(env *environment.Environment, cmd *exec.Cmd) {
	cmd.Env = env.EnvForChildCommand(env.EnvDifferencesForSubProcess())
}

func setCommandArgs(env *environment.Environment, cmd *exec.Cmd) error {

	srcPort, isGiven := env.Port.Source()

	if !isGiven {
		return nil
	}

	expression := fmt.Sprintf(`(\b%v\b)`, strconv.Itoa(srcPort))
	rgx, err := regexp.Compile(expression)

	if err != nil {
		return err
	}

	targetPortExpression := env.BackendPort()
	transformedArgs := make([]string, 0, len(cmd.Args))

	for _, arg := range cmd.Args {
		newArg := rgx.ReplaceAllLiteralString(arg, strconv.Itoa(targetPortExpression))
		transformedArgs = append(transformedArgs, newArg)
	}

	cmd.Args = transformedArgs

	return nil
}

func PipeOutputs(cmd *exec.Cmd, stdout, stderr io.Writer) error {

	cmdStdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	cmdStderr, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	go io.Copy(stdout, cmdStdout)
	go io.Copy(stderr, cmdStderr)

	return nil

}
