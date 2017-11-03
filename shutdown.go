package main

import (
	"log"
	"os"
	"os/exec"
)

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
