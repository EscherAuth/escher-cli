package environment

type Environment struct {
	Proxy Proxy
}

func New() Environment {
	return Environment{Proxy: NewProxy()}
}
