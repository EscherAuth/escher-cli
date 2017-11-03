package main

import (
	"log"

	"github.com/EscherAuth/escher/config"
	"github.com/EscherAuth/escher/keydb"
	"github.com/EscherAuth/escher/validator"
)

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
