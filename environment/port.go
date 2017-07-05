package environment

import (
	"errors"
	"net"
	"strconv"
)

type Port struct{}

func (p Port) Source() (string, bool) {
	return fetchEnv("PORT")
}

func (p Port) FindOpenAsString() (string, error) {
	port, err := findOpenPort()
	return strconv.Itoa(port), err
}

//
// 1025-65535
var NoOpenPortFoundError = errors.New("no open port found")

func findOpenPort() (int, error) {

	for i := 3000; i <= 65535; i++ {
		if portIsOpen(i) {
			return i, nil
		}
	}

	for i := 1025; i < 3000; i++ {
		if portIsOpen(i) {
			return i, nil
		}
	}

	return 0, NoOpenPortFoundError

}

func portIsOpen(port int) bool {
	var status bool
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))

	if err != nil {
		status = false
	} else {
		status = true
		conn.Close()
	}

	return status
}
