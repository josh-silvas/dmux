package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type (
	// ErrorResponse defines the error response from Nautobot which
	// includes the http.Response for additional reporting
	ErrorResponse struct {
		Response *http.Response `json:"-"`
		Detail   string         `json:"detail"`
		Errors   []message      `json:"errors"`
	}

	message struct {
		Message string `json:"message"`
	}
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

// HasError : Checks for an error message within the *http.Response object.
// This function will convert http errors into Go error data types.
func HasError(resp *http.Response) error {
	if c := resp.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	e := &ErrorResponse{
		Response: resp,
	}
	e2 := new(json.RawMessage)
	data, err := io.ReadAll(resp.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, e2); err != nil {
			return fmt.Errorf("HasError.io.ReadAll.json.Unmarshal: %w", err)
		}
		js, err := e2.MarshalJSON()
		if err != nil {
			return fmt.Errorf("HasError.io.ReadAll.json.MarshalJSON: %w", err)
		}
		e.Detail = string(js)
	}

	return e
}
