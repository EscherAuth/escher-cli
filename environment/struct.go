package environment

type Environment struct {
	Proxy Proxy
	Port  Port
}

func New() Environment {
	return Environment{
		Proxy: Proxy{},
		Port:  Port{},
	}
}
