package main

import (
	"log"
	"os"
	"os/exec"
	"time"
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

		case <-time.After(500 * time.Millisecond):

			if cmd.Process == nil {
				continue waitCycle
			}

			if cmd.ProcessState != nil {
				break waitCycle
			}

		}
	}
}
