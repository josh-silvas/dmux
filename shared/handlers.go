package shared

import (
	"os"
	"os/signal"
	"syscall"
)

// CloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean-up procedure and exiting the program.
func CloseHandler(callback func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) //nolint:govet
	go func() {
		<-c
		if callback != nil {
			callback()
		}
		os.Exit(0)
	}()
}
