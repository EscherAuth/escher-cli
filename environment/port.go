package environment

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

type Port struct{}

func (p Port) Source() (string, bool) {
	return fetchEnv("PORT")
}

//
// 1025-65535
var NoOpenPortFoundError = errors.New("no open port found")

func RequestPortFromOperationSystem() string {
	l, err := net.Listen("tcp", ":0")
	doWhatCanBeDoneWhenNoPortAvailableInAWebDevelopmentEnvironment(err)
	defer l.Close()
	address := l.Addr().String()
	parts := strings.Split(address, ":")
	return parts[len(parts)-1]
}

func doWhatCanBeDoneWhenNoPortAvailableInAWebDevelopmentEnvironment(err error) {
	if err != nil {
		// Panic (By Rincewind)
		fmt.Println(err)
		os.Exit(1)
	}
}

func (p Port) FindOpenAsString() (port string) {
	restrictedPortFromUse, sourcePortIsGiven := p.Source()

	for {
		port = RequestPortFromOperationSystem()

		if !sourcePortIsGiven {
			break
		}

		if port != restrictedPortFromUse {
			break
		}
	}

	return port
}
