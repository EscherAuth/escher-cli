package runner

func (r *runner) setEnvForCommand() error {
	env, err := r.env.EnvForChildCommand()

	if err != nil {
		return err
	}

	r.command.Env = env

	return nil
}
