package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	// NestedProvider : defines a stub circuit provider in Nautobot
	NestedProvider struct {
		CircuitCount int    `json:"circuit_count"`
		Display      string `json:"display"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Slug         string `json:"slug"`
		URL          string `json:"url"`
	}

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
		Tags         []NestedTag    `json:"tags"`
		NotesURL     string         `json:"notes_url"`
		CustomFields map[string]any `json:"custom_fields"`
	}
)

// GetProviders : method for API endpoint /api/circuits/providers/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_providers_list
func (c *Client) GetProviders(query *url.Values) ([]Provider, error) {
	req, err := c.Request(
		http.MethodGet,
		"circuits/providers/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Provider, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
