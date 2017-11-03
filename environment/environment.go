package environment

type Environment struct {
	Proxy Proxy
	Port  Port

	envDifferencesForSubProcess EnvDiff
	forwardProxyAddr            string
}

func New() *Environment {
	return &Environment{
		Proxy: Proxy{},
		Port:  Port{},
	}
}
