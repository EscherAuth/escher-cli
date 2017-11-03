package main

import (
	"log"
	"os"

	"github.com/EscherAuth/escher-cli/environment"
)

func init() {
	if len(os.Args) == 1 {
		log.Fatal("no argument given, please provide a command to execute in configured escher environment")
	}
}

func main() {
	shutdownSignals := SubscribeToShutdownSignals()
	env := environment.New()
	StartReverseProxy(env)
	StartForwardProxy(env)
	cmd := cmdBy(env)
	go RunCMD(env, cmd)
	WaitForShutdown(cmd, shutdownSignals)
}
