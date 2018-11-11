package signal_test // import "github.com/davidwalter0/tools/x/signal"

// Example to demonstrate creating and calling a signal handler
import (
	"fmt"
	"log"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/davidwalter0/tools/x/signal"
)

// create a handler example callback, called on signal by handler, log
// the signal handled and the action
func handler() (gs *signal.Handler) {
	gs = signal.NewHandler("Example callback for signal handler", nil)
	gs.Handler = func() {
		log.Printf("\n\n* %s\n", gs.Message)
		log.Printf("%s: handling signal %v shuting down in\n", gs.Name, gs.Handled)
		for i := 3; i > -1; i-- {
			fmt.Printf("%d\n", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
	}
	return gs
}

func ExampleHandler() {

	var gs = handler()
	go func() {
		// signal self
		for i := 0; i < 3; i++ {
			fmt.Printf(". ")
			time.Sleep(1)
		}
		fmt.Println()
		syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()
	gs.Wait()
}

func TestHandler(t *testing.T) {
	ExampleHandler()
}
