package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type (
	// ConsolePort : Structure to represent the console-port information derived
	// from /dcim/console-connections/
	ConsolePort struct {
		ID           string                 `json:"id"`
		URL          string                 `json:"url"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Device       struct {
			ID          string `json:"id"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
		} `json:"device"`
		Name                  string            `json:"name"`
		Label                 string            `json:"label"`
		Description           string            `json:"description"`
		ConnectedEndpointType string            `json:"connected_endpoint_type"`
		ConnectedEndpoint     ConnectedEndpoint `json:"connected_endpoint"`
		Cable                 struct {
			ID    string `json:"id"`
			URL   string `json:"url"`
			Label string `json:"label"`
		} `json:"cable"`
		Tags []Tag `json:"tags"`
	}

	// ConnectedEndpoint : Struct representing the far-side of the console connection.
	ConnectedEndpoint struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		Device struct {
			ID          string `json:"id"`
			URL         string `json:"url"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
		} `json:"device"`
		Name  string `json:"name"`
		Cable string `json:"cable"`
		Port  int
	}
)

// ConsoleConnections : Method used to fetch a list of console connections for a device.
func (c *Client) ConsoleConnections(query *url.Values) ([]ConsolePort, error) {
	req, err := c.Request(
		http.MethodGet,
		"api/dcim/console-connections/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]ConsolePort, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if e := json.Unmarshal(resp.Results, &ret); e != nil {
		return nil, fmt.Errorf("error decoding results field: %w", e)
	}

	for k, v := range ret {
		connectedEndpoint := strings.ReplaceAll(v.ConnectedEndpoint.Name, " ", "")
		if p, err := strconv.Atoi(strings.Replace(connectedEndpoint, "c", "", -1)); err == nil {
			ret[k].ConnectedEndpoint.Port = p
		}
	}

	return ret, nil
}
