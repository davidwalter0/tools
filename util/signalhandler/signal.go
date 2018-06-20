package signalhandler // import "github.com/davidwalter0/tools/util/signalhandler"

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// GracefulShutdown watch and manage
type GracefulShutdown struct {
	Watch    chan os.Signal
	Message  string
	Name     string
	Graceful func()
	Handled  os.Signal
}

// Shutdown only log info shutting down
func (gs *GracefulShutdown) Shutdown() {
	if gs.Graceful == nil {
		gs.DefaultShutdown()
	} else {
		gs.Graceful()
	}
	os.Exit(0)
}

// DefaultShutdown only log info shutting down
func (gs *GracefulShutdown) DefaultShutdown() {
	log.Printf("%s\n", gs.Message)
	log.Printf("%s: handling signal %v shutting down in ", gs.Name, gs.Handled)
	for i := 3; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("0.")
}

// PipeHandlerSignals includes watching for QUIT, PIPE and STOP
var PipeHandlerSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGHUP,
	syscall.SIGKILL,
	syscall.SIGQUIT,
	syscall.SIGSTOP,
	syscall.SIGPIPE,
}

// DefaultHandlerSignals INT, TERM, HUP
var DefaultHandlerSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGHUP,
}

// NewDefaultGracefulShutdown sets up a signal handler
//
// If no signals are given in the argument list, the default set of
// INT, TERM, HUP
func NewDefaultGracefulShutdown(text string, handleSigs ...os.Signal) *GracefulShutdown {
	if len(handleSigs) == 0 {
		handleSigs = DefaultHandlerSignals
	}
	array := strings.Split(os.Args[0], "/")
	var name = array[len(array)-1]
	var gs = &GracefulShutdown{
		Watch:    make(chan os.Signal, 1),
		Message:  fmt.Sprintf("%s: %s\n", name, text),
		Graceful: nil,
		Name:     name,
	}

	signal.Notify(gs.Watch, handleSigs...)
	return gs
}

// NewGracefulShutdown sets up a signal handler
//
// If no signals are given in the argument list, the default set of
// INT, TERM, HUP
func NewGracefulShutdown(text string, shutdown func(), handleSigs ...os.Signal) *GracefulShutdown {
	if len(handleSigs) == 0 {
		handleSigs = DefaultHandlerSignals
	}
	array := strings.Split(os.Args[0], "/")
	var name = array[len(array)-1]
	var gs = &GracefulShutdown{
		Watch:    make(chan os.Signal, 1),
		Message:  fmt.Sprintf("%s: %s\n", name, text),
		Graceful: shutdown,
		Name:     name,
	}

	signal.Notify(gs.Watch, handleSigs...)
	return gs
}

// Wait graceful shutdown callback func
func (gs *GracefulShutdown) Wait() {
	select {
	case signal := <-gs.Watch:
		gs.Handled = signal
		gs.Shutdown()
	}
}
