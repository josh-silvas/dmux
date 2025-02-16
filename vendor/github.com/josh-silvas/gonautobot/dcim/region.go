package dcim

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
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
		Parent       *Region        `json:"parent"`
		SiteCount    int            `json:"site_count"`
		Slug         string         `json:"slug"`
		URL          string         `json:"url"`
	}
)

// RegionGet : Go function to process requests for the endpoint: /api/dcim/regions/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_retrieve
func (c *Client) RegionGet(uuid string) (*Region, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/regions/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(Region)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// RegionFilter : Go function to process requests for the endpoint: /api/dcim/regions/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_list
func (c *Client) RegionFilter(q *url.Values) ([]Region, error) {
	resp := make([]Region, 0)
	if err := core.Paginate[Region](c.Client, "dcim/regions/", q, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
