package main

import (
	"os"
	"os/signal"
	"syscall"
)

var followedSignals = []os.Signal{
	syscall.SIGABRT,
	syscall.SIGALRM,
	syscall.SIGBUS,
	syscall.SIGFPE,
	syscall.SIGHUP,
	syscall.SIGILL,
	syscall.SIGINT,
	syscall.SIGKILL,
	syscall.SIGPIPE,
	syscall.SIGQUIT,
	syscall.SIGSEGV,
	syscall.SIGTERM,
	syscall.SIGTRAP,
}

// func init() {
// 	if runtime.GOOS != "windows" {
// 		followedSignals = append(
// 			followedSignals,
// 			syscall.SIGUSR1,
// 			syscall.SIGUSR2,
// 			syscall.SIGVTALRM,
// 			syscall.SIGWINCH,
// 			syscall.SIGXCPU,
// 			syscall.SIGXFSZ,
// 			syscall.SIGCHLD,
// 			syscall.SIGCONT,
// 			syscall.SIGIO,
// 			syscall.SIGCHLD,
// 			syscall.SIGCONT,
// 			syscall.SIGIO,
// 			syscall.SIGIOT,
// 			syscall.SIGPROF,
// 			syscall.SIGSTOP,
// 			syscall.SIGSYS,
// 			syscall.SIGTSTP,
// 			syscall.SIGTTIN,
// 			syscall.SIGURG,
// 			syscall.SIGTTOU)
//
// 	}
// }

func SubscribeToShutdownSignals() chan os.Signal {
	shutdownSignals := make(chan os.Signal)
	signal.Notify(shutdownSignals, followedSignals...)
	return shutdownSignals
}
