package main

import (
	"log"

	"github.com/EscherAuth/escher-cli/reverseproxy"
	"github.com/EscherAuth/escher/config"
	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
)

func main() {
	rp := reverseproxy.New(4004, NewValidator(), NewKeyDB())

	err := rp.ListenAndServeOnPort(4000)
	if err != nil {
		log.Fatal(err)
	}
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
	return keydb.NewBySlice([][2]string{[2]string{"AKIDEXAMPLE", "wJalrXUtnFEMI/K7MDENG+bPxRfiCYEXAMPLEKEY"}})
}

func NewValidator() validator.Validator {
	return validator.New(NewConfig())
}
