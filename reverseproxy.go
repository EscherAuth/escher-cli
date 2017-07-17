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
	return config.Config{
		VendorKey:       "AWS4",
		AlgoPrefix:      "AWS4",
		CredentialScope: "us-east-1/host/aws4_request",
		AuthHeaderName:  "Authorization",
		DateHeaderName:  "Date",
	}
}

func NewKeyDB() keydb.KeyDB {
	return keydb.NewByKeyValuePair("AKIDEXAMPLE", "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY")
}
