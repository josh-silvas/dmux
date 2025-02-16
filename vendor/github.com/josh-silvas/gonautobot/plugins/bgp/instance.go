package bgp

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
	nautobot "github.com/josh-silvas/gonautobot/extras"
	nautobot2 "github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// RoutingInstance : Routing Instance data representation in Nautobot.
	RoutingInstance struct {
		AutonomousSystem *AutonomousSystem        `json:"autonomous_system"`
		Created          string                   `json:"created"`
		CustomFields     map[string]interface{}   `json:"custom_fields"`
		Description      string                   `json:"description"`
		Device           *nested.Device           `json:"device"`
		Display          string                   `json:"display"`
		Endpoints        []nested.BGPPeerEndpoint `json:"endpoints"`
		ExtraAttributes  map[string]interface{}   `json:"extra_attributes"`
		ID               string                   `json:"id"`
		LastUpdated      string                   `json:"last_updated"`
		RouterID         *nested.IPAddress        `json:"router_id"`
		Status           *nautobot2.LabelValue    `json:"status"`
		Tags             []nautobot.Tag           `json:"tags"`
		URL              string                   `json:"url"`
	}
)

// BGPRoutingInstanceGet : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/:id/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_list
func (c *Client) BGPRoutingInstanceGet(uuid string) (*RoutingInstance, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("plugins/bgp/routing-instances/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(RoutingInstance)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// BGPRoutingInstanceFilter : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_retrieve
func (c *Client) BGPRoutingInstanceFilter(q *url.Values) ([]RoutingInstance, error) {
	resp := make([]RoutingInstance, 0)
	return resp, core.Paginate[RoutingInstance](c.Client, "plugins/bgp/routing-instances/", q, &resp)
}
