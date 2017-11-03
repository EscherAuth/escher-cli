package main

import (
	"log"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/EscherAuth/escher-cli/reverseproxy"
)

func StartReverseProxy(env *environment.Environment) {
	port, isGiven := env.Port.Source()

	if !isGiven {
		return
	}

	backendPort := env.BackendPort()
	reverseProxy := NewReverseProxy(backendPort)

	go func() {

		err := reverseProxy.ListenAndServe(port)

		if err != nil {
			log.Fatal(err)
		}

	}()
}

func NewReverseProxy(backendPort int) reverseproxy.ReverseProxy {
	return reverseproxy.New(backendPort, NewValidator(), NewKeyDB())
}
