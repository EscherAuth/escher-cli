package environment_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher-cli/environment/testing"
)

func TestSourcePortSetAndFound(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "3000")()

	env := environment.New()

	port, found := env.Port.Source()

	if !found {
		t.Fatal("expected port value not found")
	}

	if port != "3000" {
		t.Fatal("not expected port value: " + port)
	}

}

func TestSourcePortNotSetInTheEnv(t *testing.T) {

	env := environment.New()

	_, found := env.Port.Source()

	if found {
		t.Fatal("expected port value to be missing")
	}

}

func TestRequestPortFromOperationSystem(t *testing.T) {

	samplingTimes := 10
	ports := make([]string, 0, samplingTimes)

testTimes:
	for i := 0; i < samplingTimes; i++ {
		port := environment.RequestPortFromOperationSystem()
		fatalIfPortIsAlreadyInUse(t, port)

		for _, storedPort := range ports {
			if storedPort == port {
				continue testTimes
			}
		}

		ports = append(ports, port)
	}

	if 0 == len(ports) {
		t.Fatal("No open port found somehow!")
	}

}

func TestFindOpenAsStringPortDefaultStartingNotTaken(t *testing.T) {
	env := environment.New()
	port := env.Port.FindOpenAsString()
	fatalIfPortIsAlreadyInUse(t, port)
}

func TestFindOpenAsStringPortDefaultStartingTakenByTheHTTPPort(t *testing.T) {

	takenPort := incrementPortNumberBy(t, environment.RequestPortFromOperationSystem(), 5)
	defer SetEnvForTheTest(t, "PORT", takenPort)()
	env := environment.New()

	for i := 0; i < 100; i++ {
		port := env.Port.FindOpenAsString()
		fatalIfPortIsAlreadyInUse(t, port)
		if port == takenPort {
			t.Fatal("PORT env variable value is restricted!")
		}
	}

}

// HELPERS

func fatalIfPortIsAlreadyInUse(t testing.TB, port string) {

	_, err := strconv.Atoi(port)

	if err != nil {
		t.Fatal(err)
	}

	conn, err := net.Dial("tcp", ":"+port)

	if err == nil {
		conn.Close()
		t.Fatal("port shouldn't listen!")
	}

}

func incrementPortNumberBy(t testing.TB, port string, numberToAdd int) string {
	num, err := strconv.Atoi(port)

	if err != nil {
		t.Fatal(err)
	}

	return strconv.Itoa(num + numberToAdd)
}
