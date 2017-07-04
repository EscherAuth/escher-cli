package runner

import (
	"github.com/EscherAuth/escher-cli/environment"
)

func (r *runner) setEnvForCommand() {
	r.command.Env = environment.EnvForChildCommand(r.env)
}
