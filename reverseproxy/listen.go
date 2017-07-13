package reverseproxy

import (
	"net/http"
	"strconv"
)

func (rp *reverseProxy) ListenAndServeOnPort(port int) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rp.HandleWithValidation)
	return http.ListenAndServe(":"+strconv.Itoa(port), mux)
}
