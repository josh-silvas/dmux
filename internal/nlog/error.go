package nlog

import (
	"fmt"
)

// HandleError : error type for the Handler
type HandleError struct {
	Err error
	Msg []string
}

// Error : error message for the Handler
func (e HandleError) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Err)
}

func NewHandleError(err error, msg ...string) HandleError {
	return HandleError{Err: err, Msg: msg}
}
