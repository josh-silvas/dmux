package nautobotv1

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	// define common expected Nautobot error strings
	errDetailAuthNotProvided = "Authentication credentials were not provided."
	errDetailInvalidToken    = "Invalid token"
	errDetailNotFound        = "Not Found."
	// partial string match on raw []byte
	errRawUnknownFilter = "Unknown filter field"
)

type (
	// ErrorResponse defines the error response from Nautobot which
	// includes the http.Response for additional reporting
	ErrorResponse struct {
		Response *http.Response `json:"-"`
		Detail   string         `json:"detail"`
		// Errors field is returned by the graphql endpoint
		Errors []errorMessage `json:"errors"`
	}

	errorMessage struct {
		Message string `json:"message"`
	}
)

// define well-known Nautobot errors here for compatibility
// with errors.Is() client comparisons
var (
	// ErrAuthNotProvided : 403 response for missing authentication
	ErrAuthNotProvided = errors.New("authentication not provided")
	// ErrInvalidToken : 403 response for invalid token
	ErrInvalidToken = errors.New("invalid authentication token")
	// ErrItemNotFound : 404 response
	ErrItemNotFound = errors.New("item not found")
	// ErrUnknownFilter : 400 response for bad query parameter
	ErrUnknownFilter = errors.New("unknown query filter")
)

// Error : satisfies the 'error' interface requirements
func (e *ErrorResponse) Error() string {
	errMsg := e.Detail
	if len(e.Errors) != 0 {
		errMsg = e.joinErrors()
	}
	return fmt.Sprintf(
		"%v %v: %v",
		e.Response.Status,
		e.Response.Request.URL.String(),
		errMsg,
	)
}

// joinErrors: helper function to join array of error messages
func (e *ErrorResponse) joinErrors() string {
	messages := make([]string, 0)
	for _, err := range e.Errors {
		messages = append(messages, err.Message)
	}

	return strings.Join(messages, ", ")
}
