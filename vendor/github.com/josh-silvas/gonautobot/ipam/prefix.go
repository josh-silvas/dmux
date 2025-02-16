package ipam

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// Prefix : defines an IPAM Prefix entry in Nautobot
	Prefix struct {
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		Family       *shared.LabelValueInt  `json:"family"`
		Group        *nested.SecretsGroup   `json:"group"`
		ID           string                 `json:"id"`
		IsPool       bool                   `json:"is_pool"`
		LastUpdated  string                 `json:"last_updated"`
		Location     *nested.Location       `json:"location"`
		Name         string                 `json:"name"`
		NotesURL     string                 `json:"notes_url"`
		Prefix       string                 `json:"prefix"`
		Role         *nested.Role           `json:"role"`
		Site         *nested.Site           `json:"site"`
		Status       *shared.LabelValue     `json:"status"`
		Tags         []extras.Tag           `json:"tags"`
		Tenant       *nested.Tenant         `json:"tenant"`
		URL          string                 `json:"url"`
		VLAN         *nested.VLAN           `json:"vlan"`
		VRF          *nested.VRF            `json:"vrf"`
	}

	// PrefixAvailableIP : stub IPAddress entry returned by the /prefixes/:id/available-ips/ methods
	PrefixAvailableIP struct {
		Address string      `json:"address"`
		Family  int         `json:"family"`
		VRF     *nested.VRF `json:"vrf"`
	}

	// PrefixAvailablePrefix : stub Prefix entry returned by the /prefixes/:id/available-prefixes/ methods
	PrefixAvailablePrefix struct {
		Family int         `json:"family"`
		Prefix string      `json:"prefix"`
		VRF    *nested.VRF `json:"vrf"`
	}

	// NewPrefix : The data structure required to create a new Prefix object in Nautobot.
	NewPrefix struct {
		Description string       `json:"description"`
		ID          string       `json:"id"`
		IsPool      bool         `json:"is_pool"`
		Prefix      string       `json:"prefix"`
		Status      Status       `json:"status"`
		Tags        []extras.Tag `json:"tags"`
		VRF         *nested.VRF  `json:"vrf"`
	}

	// Status is a required parameter defining the status of the prefix ( enum)
	Status string

	// PrefixUpdate : defines writable fields to update a Prefix entry in nautobot
	PrefixUpdate struct {
		Prefix      string       `json:"prefix"`
		IsPool      bool         `json:"is_pool"`
		Tags        []extras.Tag `json:"tags"`
		Description string       `json:"description"`
		VRF         *nested.VRF  `json:"vrf"`
		Site        string       `json:"site"`
		Location    string       `json:"location"`
		Tenant      string       `json:"tenant"`
		VLAN        string       `json:"vlan"`
	}

	// PrefixesAvailablePrefixesCreateRequest : required parameters for PrefixesAvailablePrefixesCreate
	PrefixesAvailablePrefixesCreateRequest struct {
		PrefixLength int    `json:"prefix_length"`
		Status       Status `json:"status"`
	}
)

// PrefixGet : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_retrieve
func (c *Client) PrefixGet(uuid string) (*Prefix, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/prefixes/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(Prefix)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// PrefixFilter : Go function to process requests for the endpoint: /api/ipam/prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_list
func (c *Client) PrefixFilter(q *url.Values) ([]Prefix, error) {
	resp := make([]Prefix, 0)
	return resp, core.Paginate[Prefix](c.Client, "ipam/prefixes/", q, &resp)
}

// PrefixCreate : Creates a new prefix using the NewPrefix data type.
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_create
func (c *Client) PrefixCreate(prefix *NewPrefix) (*Prefix, error) {
	req, err := c.Request(http.MethodPost, "ipam/prefixes/", prefix, nil)
	if err != nil {
		return nil, err
	}

	var r Prefix
	err = c.UnmarshalDo(req, &r)
	if err != nil {
		return nil, fmt.Errorf("CreatePrefix.error.UnmarshalDo(%w)", err)
	}

	return &r, nil
}

// PrefixUpdate : Updates a Nautobot prefix by UUID
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_partial_update
func (c *Client) PrefixUpdate(uuid string, prefix *PrefixUpdate) (*Prefix, error) {
	req, err := c.Request(http.MethodPatch, fmt.Sprintf("ipam/prefixes/%s/", uuid), prefix, nil)
	if err != nil {
		return nil, err
	}

	var r Prefix
	if err := c.UnmarshalDo(req, &r); err != nil {
		return nil, fmt.Errorf("UpdatePrefix.error.UnmarshalDo(%w)", err)
	}

	return &r, nil
}

// PrefixDelete : Removes a prefix from the Nautobot DB.
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_destroy
func (c *Client) PrefixDelete(prefixID string) error {
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("ipam/prefixes/%s/", prefixID), nil, nil)
	if err != nil {
		return err
	}

	if err := c.UnmarshalDo(req, nil); err != nil {
		return fmt.Errorf("DeletePrefix.error.UnmarshalDo(%w)", err)
	}
	return nil
}

// GetPrefixAvailableIPs : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/available-ips/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_ips_list
func (c *Client) GetPrefixAvailableIPs(uuid string, q *url.Values) ([]PrefixAvailableIP, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/prefixes/%s/available-ips/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return nil, err
	}

	ret := make([]PrefixAvailableIP, 0)
	if err := c.UnmarshalDo(req, &ret); err != nil {
		return nil, fmt.Errorf("GetPrefixAvailableIPs.error.UnmarshalDo(%w)", err)
	}
	return ret, nil
}

// GetPrefixAvailablePrefixes : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/available-prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_prefixes_list
func (c *Client) GetPrefixAvailablePrefixes(uuid string, q *url.Values) ([]PrefixAvailablePrefix, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/prefixes/%s/available-prefixes/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return nil, err
	}

	ret := make([]PrefixAvailablePrefix, 0)
	if err := c.UnmarshalDo(req, &ret); err != nil {
		return nil, fmt.Errorf("GetPrefixAvailablePrefixes.error.UnmarshalDo(%w)", err)
	}
	return ret, nil
}
