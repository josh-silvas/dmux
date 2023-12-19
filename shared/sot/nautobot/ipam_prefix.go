package nautobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type (
	// Prefix : defines an IPAM Prefix entry in Nautobot
	Prefix struct {
		ID           string                 `json:"id"`
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		Family       *LabelValueInt         `json:"family"`
		IsPool       bool                   `json:"is_pool"`
		LastUpdated  string                 `json:"last_updated"`
		Location     *NestedLocation        `json:"location"`
		Name         string                 `json:"name"`
		NotesURL     string                 `json:"notes_url"`
		Prefix       string                 `json:"prefix"`
		Role         *NestedIPAMRole        `json:"role"`
		Site         *NestedSite            `json:"site"`
		Status       *LabelValue            `json:"status"`
		Tags         []NestedTag            `json:"tags"`
		Tenant       *NestedTenant          `json:"tenant"`
		URL          string                 `json:"url"`
		VLAN         *NestedVLAN            `json:"vlan"`
		VRF          *NestedVRF             `json:"vrf"`
	}

	// PrefixAvailableIP : stub IPAddress entry returned by the /prefixes/:id/available-ips/ methods
	PrefixAvailableIP struct {
		Address string     `json:"address"`
		Family  int        `json:"family"`
		VRF     *NestedVRF `json:"vrf"`
	}

	// PrefixAvailablePrefix : stub Prefix entry returned by the /prefixes/:id/available-prefixes/ methods
	PrefixAvailablePrefix struct {
		Family int        `json:"family"`
		Prefix string     `json:"prefix"`
		VRF    *NestedVRF `json:"vrf"`
	}

	// NestedPrefix : defines a stub prefix entry in Nautobot
	NestedPrefix struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	// PrefixCreate : defines structure for required request body to create an IPAM Prefix entry in nautobot
	PrefixCreate struct {
		ID          string      `json:"id"`
		Prefix      string      `json:"prefix"`
		IsPool      bool        `json:"is_pool"`
		Tags        []NestedTag `json:"tags"`
		Description string      `json:"description"`
		VRF         string      `json:"vrf"`
	}

	// PrefixUpdate : defines writable fields to update a Prefix entry in nautobot
	PrefixUpdate struct {
		Prefix      string      `json:"prefix"`
		IsPool      bool        `json:"is_pool"`
		Tags        []NestedTag `json:"tags"`
		Description string      `json:"description"`
		VRF         string      `json:"vrf"`
		Site        string      `json:"site"`
		Location    string      `json:"location"`
		Tenant      string      `json:"tenant"`
		VLAN        string      `json:"vlan"`
	}

	// PrefixLengthRequest : required parameter for PrefixesAvailablePrefixesCreate
	PrefixLengthRequest struct {
		PrefixLength int `json:"prefix_length"`
	}
)

// GetPrefix : method for API endpoint /api/ipam/prefixes/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_retrieve
func (c *Client) GetPrefix(uuid string, query *url.Values) (*Prefix, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("ipam/prefixes/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(Prefix)
	err = c.Do(req, ret)
	return ret, err
}

// GetPrefixAvailableIPs : method for API endpoint /api/ipam/prefixes/:id/available-ips/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_ips_list
func (c *Client) GetPrefixAvailableIPs(uuid string, query *url.Values) ([]PrefixAvailableIP, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("ipam/prefixes/%s/available-ips/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := make([]PrefixAvailableIP, 0)
	err = c.Do(req, &ret)
	return ret, err
}

// GetPrefixAvailablePrefixes : method for API endpoint /api/ipam/prefixes/:id/available-prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_prefixes_list
func (c *Client) GetPrefixAvailablePrefixes(uuid string, query *url.Values) ([]PrefixAvailablePrefix, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("ipam/prefixes/%s/available-prefixes/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := make([]PrefixAvailablePrefix, 0)
	err = c.Do(req, &ret)
	return ret, err
}

// PrefixesAvailablePrefixesCreate : method for API endpoint /api/ipam/prefixes/:id/available-prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_prefixes_create
func (c *Client) PrefixesAvailablePrefixesCreate(parentPrefixID string, prefixLengthReq PrefixLengthRequest) (*Prefix, bool, error) {

	req, err := c.Request(
		http.MethodPost,
		fmt.Sprintf("ipam/prefixes/%s/available-prefixes/", url.PathEscape(parentPrefixID)),
		prefixLengthReq,
		nil,
	)
	if err != nil {
		return nil, false, err
	}

	var result Prefix

	resp, err := c.RawDo(req)
	if err != nil {
		return nil, false, err
	}

	// StatusCode 204 indicates a full subnet for the parentPrefixID
	if resp.StatusCode == 204 {
		return nil, true, nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	decErr := json.NewDecoder(resp.Body).Decode(&result)
	if errors.Is(decErr, io.EOF) {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		return nil, false, fmt.Errorf("failure decoding payload: %w", decErr)
	}

	return &result, false, nil
}

// GetPrefixes : method for API endpoint /api/ipam/prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_list
func (c *Client) GetPrefixes(query *url.Values) ([]Prefix, error) {
	req, err := c.Request(
		http.MethodGet,
		"ipam/prefixes/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Prefix, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}

// PrefixesCreate creates a prefix
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_create
func (c *Client) PrefixesCreate(prefix *PrefixCreate) (*Prefix, error) {

	for _, tag := range prefix.Tags {
		_, err := c.TagUpsert(tag)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.Request(http.MethodPost, "ipam/prefixes/", prefix, nil)
	if err != nil {
		return nil, err
	}

	var result Prefix
	err = c.Do(req, &result)
	if err != nil {
		return nil, fmt.Errorf("return code error: %w", err)
	}

	return &result, nil
}

// PrefixesUpdate updates a prefix by prefix ID
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_partial_update
func (c *Client) PrefixesUpdate(prefixID string, prefix *PrefixUpdate) (*Prefix, error) {

	for _, tag := range prefix.Tags {
		_, err := c.TagUpsert(tag)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.Request(http.MethodPatch, fmt.Sprintf("ipam/prefixes/%s/", prefixID), prefix, nil)

	if err != nil {
		return nil, err
	}

	var result Prefix

	err = c.Do(req, &result)
	if err != nil {
		return nil, fmt.Errorf("return code error: %w", err)
	}

	return &result, nil
}

// PrefixesDelete deletes a prefix by prefix ID
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_destroy
func (c *Client) PrefixesDelete(prefixID string) error {

	req, err := c.Request(http.MethodDelete, fmt.Sprintf("ipam/prefixes/%s/", prefixID), nil, nil)

	if err != nil {
		return err
	}

	err = c.Do(req, nil)
	if err != nil {
		return fmt.Errorf("return code error: %w", err)
	}

	return nil
}
