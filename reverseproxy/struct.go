package reverseproxy

import (
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
)

type ReverseProxy interface {
	ListenAndServeOnPort(port int) error
}

type reverseProxy struct {
	reverseProxy *httputil.ReverseProxy
	validator    validator.Validator
	keyDB        keydb.KeyDB
}

func New(backendPort int, validator validator.Validator, keyDB keydb.KeyDB) ReverseProxy {
	return &reverseProxy{
		keyDB:        keyDB,
		validator:    validator,
		reverseProxy: httputil.NewSingleHostReverseProxy(backendServerURLBy(backendPort)),
	}
}

func backendServerURLBy(port int) *url.URL {
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = ":" + strconv.Itoa(port)
	return u
}
