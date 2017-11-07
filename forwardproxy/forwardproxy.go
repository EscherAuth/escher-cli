package forwardproxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/EscherAuth/escher/request"
	"github.com/EscherAuth/escher/signer"
)

func New(eSigner signer.Signer) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: newDirector(eSigner)}
}

func d(req *http.Request) {
	if os.Getenv("ESCHER_DEBUG") == "true" {
		fmt.Println("---------------")
		bs, _ := httputil.DumpRequest(req, false)
		fmt.Println(string(bs))
		fmt.Println("- - - - - - - -")
	}
}
func newDirector(eSigner signer.Signer) func(req *http.Request) {
	return func(req *http.Request) {

		err := signRequest(eSigner, req)

		fmt.Println(req)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func l() { fmt.Println(time.Now()) }

func signRequest(eSigner signer.Signer, req *http.Request) error {
	escherRequest, err := request.NewFromHTTPRequest(req)

	if err != nil {
		return err
	}

	signedRequest, err := eSigner.SignRequest(escherRequest, []string{})

	if err != nil {
		return err
	}

	err = signedRequest.UpdateHTTPRequest(req)

	if err != nil {
		return err
	}

	return nil
}
