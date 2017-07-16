package reverseproxy

import (
	"net/http"

	"github.com/EscherAuth/escher/request"
)

func (rp *reverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	apiKey, err := rp.validator.Validate(request.NewFromHTTPRequest(r), rp.keyDB, nil)

	if err != nil {
		w.Header().Set("WWW-Authenticate", "EscherAuth")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	r.Header.Set("X-Escher-Key", apiKey)

	rp.reverseProxy.ServeHTTP(w, r)

}
