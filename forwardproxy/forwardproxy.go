package forwardproxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/EscherAuth/escher/request"
	"github.com/EscherAuth/escher/signer"
)

func New(eSigner signer.Signer) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: newDirector(eSigner)}
}

func newDirector(eSigner signer.Signer) func(req *http.Request) {
	return func(req *http.Request) {

		err := signRequest(eSigner, req)

		if err != nil {
			fmt.Println(err)
		}

	}
}

func signRequest(eSigner signer.Signer, req *http.Request) error {
	escherRequest, err := request.NewFromHTTPRequest(req)

	if err != nil {
		return err
	}

	signedRequest, err := eSigner.SignRequest(escherRequest, headersToSignBy(req))

	if err != nil {
		return err
	}

	err = signedRequest.UpdateHTTPRequest(req)

	if err != nil {
		return err
	}

	return nil
}

// Hop-by-hop headers. These are removed when sent to the backend.
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var hopHeaders = map[string]struct{}{
	"connection":          struct{}{},
	"proxy-connection":    struct{}{}, // non-standard but still sent by libcurl and rejected by e.g. google
	"keep-alive":          struct{}{},
	"proxy-authenticate":  struct{}{},
	"proxy-authorization": struct{}{},
	"te":                struct{}{}, // canonicalized version of "TE"
	"trailer":           struct{}{}, // not Trailers per URL above; http://www.rfc-editor.org/errata_search.php?eid=4522
	"transfer-encoding": struct{}{},
	"upgrade":           struct{}{},
}

func headersToSignBy(req *http.Request) []string {
	headerNames := make([]string, 0, len(req.Header))

	for name, _ := range req.Header {
		if _, ok := hopHeaders[strings.ToLower(name)]; !ok {
			headerNames = append(headerNames, name)
		}
	}

	return headerNames
}
