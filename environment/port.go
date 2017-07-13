package environment

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Port struct{}

func (p Port) Source() (int, bool) {
	portAsString, isGiven := fetchEnv("PORT")

	if !isGiven {
		return 0, isGiven
	}

	port, err := strconv.Atoi(portAsString)
	return port, err == nil
}

//
// 1025-65535
var NoOpenPortFoundError = errors.New("no open port found")

func RequestPortFromOperationSystem() int {
	l, err := net.Listen("tcp", ":0")
	doWhatCanBeDoneWhenNoPortAvailableInAWebDevelopmentEnvironment(err)
	defer l.Close()
	address := l.Addr().String()
	parts := strings.Split(address, ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	doWhatCanBeDoneWhenNoPortAvailableInAWebDevelopmentEnvironment(err)
	return port
}

func doWhatCanBeDoneWhenNoPortAvailableInAWebDevelopmentEnvironment(err error) {
	if err != nil {
		// Panic (By Rincewind)
		fmt.Println(err)
		os.Exit(1)
	}
}

func (p Port) FindOpenAsString() (port int) {
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
