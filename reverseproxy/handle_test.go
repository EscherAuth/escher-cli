package reverseproxy_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/EscherAuth/escher-cli/reverseproxy"
	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
	"github.com/EscherAuth/escher/validator/mock"
	"github.com/stretchr/testify/assert"
)

var backendServerPort int

func init() {

	diff, err := environment.New().EnvDifferencesForSubProcess()

	if err != nil {
		log.Fatal(err)
	}

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

	proxy.HandleWithValidation(response, request)

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

	proxy.HandleWithValidation(response, request)

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
	go server.ListenAndServe()

	return func() {
		err := server.Shutdown(context.Background())

		if err != nil {
			t.Fatal(err)
		}
	}
}
