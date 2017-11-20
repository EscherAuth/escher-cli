package reverseproxy_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/EscherAuth/escher-cli/reverseproxy"
	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
	"github.com/EscherAuth/escher/validator/mock"
	"github.com/stretchr/testify/assert"
)

var backendServerPort int

func init() {

	diff := environment.New().EnvDifferencesForSubProcess()

	openPortNumber, err := strconv.Atoi(diff["PORT"])

	if err != nil {
		log.Fatal(err)
	}

	backendServerPort = openPortNumber

}

func TestHandlingValidRequestForwardedToTheBackend(t *testing.T) {
	defer backendServerIsListening(t, backendServerPort)()

	validator := mock.New()
	keyDB := keydb.NewByKeyValuePair("foo", "baz")

	proxy := reverseproxy.New(backendServerPort, validator, keyDB)

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/testing", bytes.NewBuffer([]byte{}))

	if err != nil {
		t.Fatal(err)
	}

	proxy.ServeHTTP(response, request)

	actuallyBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "PATH_INFO: testing", string(actuallyBody))

}

func TestHandlingValidationErrorRespondedWithoutBotherinTheBackendService(t *testing.T) {
	defer grumpyBackendServerIsListening(t, backendServerPort)()

	validatorMock := mock.New()
	validatorMock.AddValidationResult("test", validator.InvalidEscherKey)

	keyDB := keydb.NewByKeyValuePair("foo", "baz")

	proxy := reverseproxy.New(backendServerPort, validatorMock, keyDB)

	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/testing", bytes.NewBuffer([]byte{}))

	if err != nil {
		t.Fatal(err)
	}

	proxy.ServeHTTP(response, request)

	actuallyBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Unauthorized\n", string(actuallyBody))
	assert.Equal(t, "EscherAuth", response.Header().Get("WWW-Authenticate"))
	assert.Equal(t, 401, response.Code)

}

func grumpyBackendServerIsListening(t testing.TB, port int) func() {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		t.Fail()
		fmt.Fprint(w, "no...")
	}

	return backendServerIsListeningWith(t, port, handleFunc)
}

func backendServerIsListening(t testing.TB, port int) func() {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "PATH_INFO: %v", r.URL.Path[1:])
	}

	return backendServerIsListeningWith(t, port, handleFunc)
}

func backendServerIsListeningWith(t testing.TB, port int, handleFunc func(http.ResponseWriter, *http.Request)) func() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleFunc)

	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: mux}

	WaitForPortToBeOpen(port)

	go server.ListenAndServe()

	return func() {
		err := server.Shutdown(context.Background())

		if err != nil {
			t.Fatal(err)
		}
	}
}

// Check if a port is available
func Check(port int) bool {

	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false
	}

	// close the server
	server.Close()
	time.Sleep(500 * time.Millisecond)

	// we successfully used and closed the port
	// so it's now available to be used again
	return true

}

func WaitForPortToBeOpen(port int) {
	for !Check(port) {
		time.Sleep(500 * time.Millisecond)
	}
}
