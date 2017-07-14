package reverseproxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
)

type ReverseProxy interface {
	ListenAndServeOnPort(port int) error
	HandleWithValidation(http.ResponseWriter, *http.Request)
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
		reverseProxy: httpUtilsReverseProxyBy(backendPort),
	}
}
