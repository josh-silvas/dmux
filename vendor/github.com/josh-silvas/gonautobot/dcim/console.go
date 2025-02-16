package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type (
	// ConsolePort : Data representation of a Console Port in Nautobot.
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
		Tags []extras.Tag `json:"tags"`
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

// ConsoleConnectionFilter : Method used to fetch a list of console connections for a device.
func (c *Client) ConsoleConnectionFilter(q *url.Values) ([]ConsolePort, error) {
	req, err := c.Request(http.MethodGet, "dcim/console-connections/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]ConsolePort, 0)
	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if e := json.Unmarshal(resp.Results, &ret); e != nil { // nolint: musttag
		return nil, fmt.Errorf("ConsoleConnections.error.json.Unmarshal(%w)", e)
	}

	for k, v := range ret {
		ce := strings.ReplaceAll(v.ConnectedEndpoint.Name, " ", "")
		if p, err := strconv.Atoi(strings.Replace(ce, "c", "", -1)); err == nil {
			ret[k].ConnectedEndpoint.Port = p
		}
	}

	return ret, nil
}
