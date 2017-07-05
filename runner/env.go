package runner

func (r *runner) setEnvForCommand() error {
	envChanges, err := r.env.EnvDifferencesForSubProcess()

	if err != nil {
		return err
	}

	r.envChanges = envChanges
	r.command.Env = r.env.EnvForChildCommand(r.envChanges)

	return nil
}
