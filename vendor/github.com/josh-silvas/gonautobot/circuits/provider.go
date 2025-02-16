package circuits

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"net/http"
	"net/url"
	"time"
)

type (
	// Provider : defines a circuit provider in Nautobot
	Provider struct {
		ID           string         `json:"id"`
		Display      string         `json:"display"`
		URL          string         `json:"url"`
		Name         string         `json:"name"`
		Slug         string         `json:"slug"`
		Asn          int            `json:"asn"`
		Account      string         `json:"account"`
		PortalURL    string         `json:"portal_url"`
		NocContact   string         `json:"noc_contact"`
		AdminContact string         `json:"admin_contact"`
		Comments     string         `json:"comments"`
		CircuitCount int            `json:"circuit_count"`
		Created      string         `json:"created"`
		LastUpdated  time.Time      `json:"last_updated"`
		Tags         []extras.Tag   `json:"tags"`
		NotesURL     string         `json:"notes_url"`
		CustomFields map[string]any `json:"custom_fields"`
	}
)

// ProviderFilter : Go function to process requests for the endpoint: /api/circuits/providers/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_providers_list
func (c *Client) ProviderFilter(q *url.Values) ([]Provider, error) {
	req, err := c.Request(http.MethodGet, "circuits/providers/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]Provider, 0)
	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("GetProviders.error.json.Unmarshall(%w)", err)
	}
	return ret, err
}
