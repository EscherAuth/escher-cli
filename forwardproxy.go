package main

import (
	"net/http"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/EscherAuth/escher-cli/forwardproxy"
	"github.com/EscherAuth/escher/signer"
)

func StartForwardProxy(env *environment.Environment) {
	p := forwardproxy.New(signer.New(NewConfig()))

	go http.ListenAndServe(env.ForwardProxyAddr(), p)
}
