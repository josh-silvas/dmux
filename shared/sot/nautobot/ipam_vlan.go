package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	// VLAN : defines a vlan entry in Nautobot
	VLAN struct {
		ID           string                 `json:"id"`
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		LastUpdated  string                 `json:"last_updated"`
		Location     *NestedLocation        `json:"location"`
		Name         string                 `json:"name"`
		NotesURL     string                 `json:"notes_url"`
		PrefixCount  int                    `json:"prefix_count"`
		Role         *NestedIPAMRole        `json:"role"`
		Site         *NestedSite            `json:"site"`
		Status       *LabelValue            `json:"status"`
		Tags         []NestedTag            `json:"tags"`
		Tenant       *NestedTenant          `json:"tenant"`
		URL          string                 `json:"url"`
		VID          int                    `json:"vid"`
	}

	// NestedVLAN : defines a stub vlan entry in Nautobot
	NestedVLAN struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		VID     int    `json:"vid"`
	}
)

// GetVLAN : method for API endpoint /api/ipam/vlans/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_vlans_retrieve
func (c *Client) GetVLAN(uuid string, query *url.Values) (*VLAN, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("ipam/vlans/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(VLAN)
	err = c.Do(req, ret)
	return ret, err
}

// GetVLANs : method for API endpoint /api/ipam/vlans/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_vlans_list
func (c *Client) GetVLANs(query *url.Values) ([]VLAN, error) {
	req, err := c.Request(
		http.MethodGet,
		"ipam/vlans/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]VLAN, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
