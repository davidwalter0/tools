package signal // import "github.com/davidwalter0/tools/x/signal"

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Handler watch and manage signals
type Handler struct {
	Watch   chan os.Signal
	Message string
	Name    string
	Handler func()
	Handled os.Signal
}

// Shutdown only log info shutting down
func (gs *Handler) Shutdown() {
	if gs.Handler == nil {
		gs.Default()
	} else {
		gs.Handler()
	}
	os.Exit(0)
}

// Default only log info shutting down
func (gs *Handler) Default() {
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

// DefaultSignals INT, TERM, HUP
var DefaultSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGHUP,
}

// UsrSignals INT, TERM, HUP
var UsrSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGHUP,
	syscall.SIGUSR1,
	syscall.SIGUSR2,
}

// NewDefault sets up a signal handler
//
// If no signals are given in the argument list, the default set of
// INT, TERM, HUP
func NewDefault(text string, handleSigs ...os.Signal) *Handler {
	if len(handleSigs) == 0 {
		handleSigs = DefaultSignals
	}
	array := strings.Split(os.Args[0], "/")
	var name = array[len(array)-1]
	var gs = &Handler{
		Watch:   make(chan os.Signal, 1),
		Message: fmt.Sprintf("%s: %s\n", name, text),
		Handler: nil,
		Name:    name,
	}

	signal.Notify(gs.Watch, handleSigs...)
	return gs
}

// NewUsr sets up a signal handler
//
// If no signals are given in the argument list, the default set of
// INT, TERM, HUP
func NewUsr(text string, handleSigs ...os.Signal) *Handler {
	if len(handleSigs) == 0 {
		handleSigs = UsrSignals
	}
	array := strings.Split(os.Args[0], "/")
	var name = array[len(array)-1]
	var gs = &Handler{
		Watch:   make(chan os.Signal, 1),
		Message: fmt.Sprintf("%s: %s\n", name, text),
		Handler: nil,
		Name:    name,
	}

	signal.Notify(gs.Watch, handleSigs...)
	return gs
}

// NewHandler sets up a signal handler
//
// If no signals are given in the argument list, the default set of
// INT, TERM, HUP
func NewHandler(text string, shutdown func(), handleSigs ...os.Signal) *Handler {
	if len(handleSigs) == 0 {
		handleSigs = DefaultSignals
	}
	array := strings.Split(os.Args[0], "/")
	var name = array[len(array)-1]
	var gs = &Handler{
		Watch:   make(chan os.Signal, 1),
		Message: fmt.Sprintf("%s: %s\n", name, text),
		Handler: shutdown,
		Name:    name,
	}

	signal.Notify(gs.Watch, handleSigs...)
	return gs
}

// Wait save the signal received, call signal shutdown callback func
func (gs *Handler) Wait() {
	select {
	case signal := <-gs.Watch:
		gs.Handled = signal
		gs.Shutdown()
	}
}
