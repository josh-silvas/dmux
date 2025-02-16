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
	// VLAN : defines a vlan entry in Nautobot
	VLAN struct {
		ID           string                 `json:"id"`
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		Group        *nested.SecretsGroup   `json:"group"`
		LastUpdated  string                 `json:"last_updated"`
		Location     *nested.Location       `json:"location"`
		Name         string                 `json:"name"`
		NotesURL     string                 `json:"notes_url"`
		PrefixCount  int                    `json:"prefix_count"`
		Role         *nested.Role           `json:"role"`
		Site         *nested.Site           `json:"site"`
		Status       *shared.LabelValue     `json:"status"`
		Tags         []extras.Tag           `json:"tags"`
		Tenant       *nested.Tenant         `json:"tenant"`
		URL          string                 `json:"url"`
		VID          int                    `json:"vid"`
	}
)

// VLANGet : Go function to process requests for the endpoint: /api/ipam/vlans/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_vlans_retrieve
func (c *Client) VLANGet(uuid string) (*VLAN, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/vlans/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}
	ret := new(VLAN)
	return ret, c.UnmarshalDo(req, ret)
}

// VLANFilter : Go function to process requests for the endpoint: /api/ipam/vlans/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_vlans_list
func (c *Client) VLANFilter(q *url.Values) ([]VLAN, error) {
	resp := make([]VLAN, 0)
	return resp, core.Paginate[VLAN](c.Client, "ipam/vlans/", q, &resp)
}
