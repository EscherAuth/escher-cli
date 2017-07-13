package reverseproxy

import (
	"net/http"
)

func (rp *reverseProxy) HandleWithValidation(w http.ResponseWriter, r *http.Request) {

	// apiKey, err := rp.validator.Validate(request.NewFromHTTPRequest(r), rp.keyDB, nil)

	// if err != nil {
	// 	w.Header().Set("WWW-Authenticate", "EscherAuth")
	// 	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	// 	return
	// }

	// r.Header.Set("X-Escher-Key", apiKey)

	rp.reverseProxy.ServeHTTP(w, r)

}
