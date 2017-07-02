package environment

type Environment struct {
	Proxy Proxy
	Host  string `env:"HOST"`
}

func New() Environment {
	return Environment{Proxy: NewProxy(), Host: getHost()}
}
