package reverseproxy

import (
	"net/http"
	"strconv"
)

// TODO: cover
func (rp *reverseProxy) ListenAndServe(port int) error {
	return http.ListenAndServe(":"+strconv.Itoa(port), rp)
}
