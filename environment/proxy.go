package environment

type Proxy struct{}

func (p Proxy) HTTP() (string, bool) {
	return fetchEnv("HTTP_PROXY")
}

func (p Proxy) HTTPS() (value string, found bool) {
	value, found = fetchEnv("HTTPS_PROXY")

	if !found {
		value, found = p.HTTP()
	}

	return
}
