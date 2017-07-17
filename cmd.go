package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/EscherAuth/escher-cli/command"
	"github.com/EscherAuth/escher-cli/environment"
)

func cmdBy(env *environment.Environment) *exec.Cmd {

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	err := command.ConfigureBy(env, cmd)

	if err != nil {
		log.Fatal(err)
	}

	err = command.PipeOutputs(cmd, os.Stdout, os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	return cmd

}

func RunCMD(env *environment.Environment, cmd *exec.Cmd) {

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}
