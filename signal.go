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
	syscall.SIGCHLD,
	syscall.SIGCONT,
	syscall.SIGFPE,
	syscall.SIGHUP,
	syscall.SIGILL,
	syscall.SIGINT,
	syscall.SIGIO,
	syscall.SIGIOT,
	syscall.SIGKILL,
	syscall.SIGPIPE,
	syscall.SIGPROF,
	syscall.SIGQUIT,
	syscall.SIGSEGV,
	syscall.SIGSTOP,
	syscall.SIGSYS,
	syscall.SIGTERM,
	syscall.SIGTRAP,
	syscall.SIGTSTP,
	syscall.SIGTTIN,
	syscall.SIGTTOU,
	syscall.SIGURG,
	syscall.SIGUSR1,
	syscall.SIGUSR2,
	syscall.SIGVTALRM,
	syscall.SIGWINCH,
	syscall.SIGXCPU,
	syscall.SIGXFSZ,
}

func SubscribeToShutdownSignals() chan os.Signal {
	shutdownSignals := make(chan os.Signal)
	signal.Notify(shutdownSignals, followedSignals...)
	return shutdownSignals
}
