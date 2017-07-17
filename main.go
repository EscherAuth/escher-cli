package main

import (
	"log"
	"os"
	"os/exec"

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

	StartListenInTheBackgroundWithReverseProxy(env)

	cmd := cmdBy(env)
	go RunCMD(env, cmd)
	WaitForShutdown(cmd, shutdownSignals)
}

func WaitForShutdown(cmd *exec.Cmd, shutdownSignals chan os.Signal) {
waitCycle:
	for {
		select {

		case sig := <-shutdownSignals:
			err := cmd.Process.Signal(sig)

			if err != nil {
				log.Println(err)
			}
		default:

			if cmd.Process == nil {
				continue waitCycle
			}

			if cmd.ProcessState != nil {
				break waitCycle
			}

		}
	}
}
