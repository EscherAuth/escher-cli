package environment

type Proxy struct {
	HTTP  string `env:"HTTP_PROXY"`
	HTTPS string `env:"HTTPS_PROXY"`
}

func NewProxy() Proxy {
	return Proxy{
		HTTP:  fetchEnv("HTTP_PROXY"),
		HTTPS: httpsProxyEnv(),
	}
}

func httpsProxyEnv() string {
	httpsProxy := fetchEnv("HTTPS_PROXY")

	if httpsProxy != "" {
		return httpsProxy
	}

	return fetchEnv("HTTP_PROXY")
}
