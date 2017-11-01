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

// Source is the Port where the Guarded Server originally would expect the incoming Http requests
func (p Port) Source() (int, bool) {
	portAsString, isGiven := fetchEnv("PORT")

	if !isGiven {
		return 0, isGiven
	}

	port, err := strconv.Atoi(portAsString)
	return port, err == nil
}

// NoOpenPortFoundError the error what returned when no open port found
var NoOpenPortFoundError = errors.New("no open port found")

// RequestPortFromOperationSystem get a port from between 1025-65535
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

func (p Port) FindOpen() (port int) {
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
