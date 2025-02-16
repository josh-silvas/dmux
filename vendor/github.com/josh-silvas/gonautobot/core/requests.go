package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client : Requests data type client.
type Client struct {
	Client *http.Client
	Token  string
	URL    string
}

// Paginate : Helper function to abstract out Nautobot paginated responses.
func Paginate[T any](c *Client, uri string, q *url.Values, ret *[]T) error {
	fullURL := ""
	for {
		if fullURL == "" {
			uri = strings.TrimPrefix(uri, "/")
			fullURL = fmt.Sprintf("%s%s", c.URL, uri)
		}

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fullURL, nil)
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Token))
		if q != nil && len(*q) > 0 {
			req.URL.RawQuery = q.Encode()
		}

		resp := new(shared.ResponseList)
		r := make([]T, 0)
		if err = c.UnmarshalDo(req, resp); err != nil {
			return err
		}
		if err = json.Unmarshal(resp.Results, &r); err != nil {
			return fmt.Errorf("Paginate.error.json.Unmarshal(%w)", err)
		}

		*ret = append(*ret, r...)
		if resp.Next == "" || fullURL == resp.Next {
			break
		}
		log.Info().Msgf("Paginate.next=%s; Total.items=%d", resp.Next, resp.Count)
		fullURL = resp.Next
	}
	return nil
}

// Request : Crafts an HTTP request to Nautobot.
func (c *Client) Request(method, uri string, body interface{}, params *url.Values) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(body); err != nil {
			return nil, err
		}
	}
	uri = strings.TrimPrefix(uri, "/")

	req, err := http.NewRequestWithContext(context.Background(), method, fmt.Sprintf("%s%s", c.URL, uri), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.Token)
	if params != nil && len(*params) > 0 {
		req.URL.RawQuery = params.Encode()
	}

	return req, nil
}

// Do : Performs an HTTP requests and returns the raw *http.Response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, HasError(resp)
}

// UnmarshalDo : Performs an HTTP request to the resource defined in the request.
func (c *Client) UnmarshalDo(req *http.Request, v interface{}) error {
	resp, err := c.Do(req)
	if resp != nil && resp.Body != nil {
		defer func(r *http.Response) {
			if err := r.Body.Close(); err != nil {
				log.Warn().Err(err).Msg("UnmarshalDo.resp.Body.Close()")
			}
		}(resp)
	}
	if err != nil {
		return fmt.Errorf("c.Do.error: %w", err)
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		if _, err := io.Copy(v, resp.Body); err != nil {
			return fmt.Errorf("UnmarshalDo.io.Copy(): %w", err)
		}
	default:
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return fmt.Errorf("error.UnmarshalDo.json.NewDecoder(%w)", err)
		}
	}
	return nil
}
