package reverseproxy

import (
	"net/http/httputil"
	"net/url"
	"strconv"
)

func httpUtilsReverseProxyBy(port int) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(backendServerURLBy(port))
}

func backendServerURLBy(port int) *url.URL {
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = ":" + strconv.Itoa(port)
	return u
}
