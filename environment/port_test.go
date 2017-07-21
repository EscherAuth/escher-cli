package environment_test

import (
	"strconv"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	. "github.com/EscherAuth/escher/testing/env"
)

func TestSourcePortSetAndFound(t *testing.T) {
	defer SetEnvForTheTest(t, "PORT", "3000")()

	env := environment.New()

	port, found := env.Port.Source()

	if !found {
		t.Fatal("expected port value not found")
	}

	if port != 3000 {
		t.Fatal("not expected port value: " + strconv.Itoa(port))
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
	ports := make([]int, 0, samplingTimes)

testTimes:
	for i := 0; i < samplingTimes; i++ {
		port := environment.RequestPortFromOperationSystem()
		FatalIfPortIsAlreadyInUse(t, port)

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
	FatalIfPortIsAlreadyInUse(t, port)
}

func TestFindOpenAsStringPortDefaultStartingTakenByTheHTTPPort(t *testing.T) {
	takenPort := environment.RequestPortFromOperationSystem() + 2
	defer SetEnvForTheTest(t, "PORT", strconv.Itoa(takenPort))()
	env := environment.New()

	for i := 0; i < 100; i++ {
		port := env.Port.FindOpenAsString()
		FatalIfPortIsAlreadyInUse(t, port)
		if port == takenPort {
			t.Fatal("PORT env variable value is restricted!")
		}
	}

}
