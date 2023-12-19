package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	// Region : defines a region object in Nautobot
	Region struct {
		ID           string         `json:"id"`
		Created      string         `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Depth        int            `json:"_depth"`
		Description  string         `json:"description" datastore:",noindex"`
		Display      string         `json:"display"`
		LastUpdated  string         `json:"last_updated"`
		Name         string         `json:"name"`
		NotesURL     string         `json:"notes_url"`
		Parent       *NestedRegion  `json:"parent"`
		SiteCount    int            `json:"site_count"`
		Slug         string         `json:"slug"`
		URL          string         `json:"url"`
	}

	// NestedRegion : defines a stub region entry in Nautobot
	NestedRegion struct {
		ID        string `json:"id"`
		Depth     int    `json:"_depth"`
		Display   string `json:"display"`
		Name      string `json:"name"`
		SiteCount int    `json:"site_count"`
		Slug      string `json:"slug"`
		URL       string `json:"url"`
	}
)

// GetRegion : method for API endpoint /api/dcim/regions/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_retrieve
func (c *Client) GetRegion(uuid string, query *url.Values) (*Region, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("dcim/regions/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(Region)
	err = c.Do(req, ret)
	return ret, err
}

// GetRegions : method for API endpoint /api/dcim/regions/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_list
func (c *Client) GetRegions(query *url.Values) ([]Region, error) {
	req, err := c.Request(
		http.MethodGet,
		"dcim/regions/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Region, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
