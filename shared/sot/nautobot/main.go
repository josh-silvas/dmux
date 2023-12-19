package nautobot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	// Client : Stored memory objects for the Nautobot client.
	Client struct {
		client   *http.Client
		instance *url.URL
		log      *logrus.Logger
		token    string
	}

	// LabelValue defines a repeating structure used throughout the API
	LabelValue struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}

	// LabelValueInt defines a repeating structure used throughout the API
	LabelValueInt struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	}

	rawListResponse struct {
		Count    int             `json:"count"`
		Next     string          `json:"next"`
		Previous string          `json:"previous"`
		Results  json.RawMessage `json:"results"`
	}
)

// New : Function used to create a new Nautobot client data type.
func New(token, nbURL string, opts ...Option) (*Client, error) {
	client := &Client{
		token: token,
		instance: func() *url.URL {
			//nolint:errcheck
			u, _ := url.Parse(nbURL)
			return u
		}(),
	}
	err := client.processOptions(opts...)
	return client, err
}

// Request : creates an API request; a relative URI should be provided for uri
// and should not have a leading slash for a proper url.Parse() merge
func (c *Client) Request(method, uri string, body interface{}, queryParameters *url.Values) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	u, err := c.instance.Parse(uri)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(context.Background(), method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.token)

	if queryParameters != nil && len(*queryParameters) > 0 {
		req.URL.RawQuery = queryParameters.Encode()
	}

	return req, nil
}

// checkResponse : checks for a non-2xx HTTP response code
// and attempts to unmarshal an error payload from Nautobot;
// returns custom ErrorResponse which includes original http.Response
func checkResponse(resp *http.Response) error {
	if c := resp.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	e := &ErrorResponse{
		Response: resp,
	}
	data, err := io.ReadAll(resp.Body)
	if err == nil && data != nil {
		//nolint:errcheck
		_ = json.Unmarshal(data, e)
	}

	// check for well-known errors and return pre-defined error
	switch {
	case strings.EqualFold(e.Detail, errDetailAuthNotProvided):
		return ErrAuthNotProvided
	case strings.EqualFold(e.Detail, errDetailInvalidToken):
		return ErrInvalidToken
	case strings.EqualFold(e.Detail, errDetailNotFound):
		return ErrItemNotFound
	case strings.Contains(string(data), errRawUnknownFilter):
		return ErrUnknownFilter
	}

	return e
}

// RawDo : Sends the API request and returns the raw response and any error.
// We should normally use do() which JSON-decodes and closes the response Body, but
// if there is a non-JSON endpoint or other reason to get the raw response, then
// this can be used.
func (c *Client) RawDo(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
	}
	return resp, err
}

// Do : sends an API request and JSON-decodes the API response.
// The response is stored in the value pointed to by v
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if resp != nil {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
	}

	err = checkResponse(resp)
	if err != nil {
		return err
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if errors.Is(decErr, io.EOF) {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = fmt.Errorf("failure decoding payload: %w", decErr)
		}
	}
	return err
}
