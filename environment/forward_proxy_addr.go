package environment

import (
	"fmt"
	"strconv"
	"strings"
)

func (e *Environment) ForwardProxyAddr() string {

	if e.forwardProxyAddr != "" {
		return e.forwardProxyAddr
	}

	port := e.Port.FindOpen()
	portAsString := strconv.Itoa(port)
	proxyAddress := fmt.Sprintf("localhost:%s", portAsString)

	diff := e.EnvDifferencesForSubProcess()
	for _, name := range []string{"http_proxy", "https_proxy"} {
		diff[strings.ToLower(name)] = proxyAddress
		diff[strings.ToUpper(name)] = proxyAddress
	}

	e.forwardProxyAddr = proxyAddress
	return proxyAddress

}
