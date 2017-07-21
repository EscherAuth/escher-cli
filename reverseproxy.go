package main

import (
	"log"

	"github.com/EscherAuth/escher-cli/environment"
	"github.com/EscherAuth/escher-cli/reverseproxy"
	"github.com/EscherAuth/escher/config"
	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
)

func StartListenInTheBackgroundWithReverseProxy(env *environment.Environment) {
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

func NewValidator() validator.Validator {
	return validator.New(NewConfig())
}

func NewConfig() config.Config {
	c, err := config.NewFromENV()

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func NewKeyDB() keydb.KeyDB {
	keyDB, err := keydb.NewFromENV()

	if err != nil {
		log.Fatal(err)
	}

	return keyDB
}
