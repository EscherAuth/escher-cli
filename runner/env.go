package runner

func (r *runner) setEnvForCommand() error {
	r.command.Env = r.env.EnvForChildCommand(r.env.EnvDifferencesForSubProcess())

	return nil
}
