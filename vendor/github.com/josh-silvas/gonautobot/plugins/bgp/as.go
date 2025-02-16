package bgp

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
	// AutonomousSystem : Autonomous System data struct representation.
	AutonomousSystem struct {
		ASN          int                    `json:"asn"`
		Created      string                 `json:"created"`
		CustomFields map[string]interface{} `json:"custom_fields"`
		Description  string                 `json:"description"`
		Display      string                 `json:"display"`
		ID           string                 `json:"id"`
		LastUpdated  string                 `json:"last_updated"`
		Provider     *nested.Provider       `json:"provider"`
		Status       *shared.LabelValue     `json:"status"`
		Tags         []extras.Tag           `json:"tags"`
		URL          string                 `json:"url"`
	}
)

// BGPAutonomousSystemGet : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/:id/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_list
func (c *Client) BGPAutonomousSystemGet(uuid string) (*AutonomousSystem, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("plugins/bgp/autonomous-systems/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(AutonomousSystem)
	return ret, c.UnmarshalDo(req, ret)
}

// BGPAutonomousSystemFilter : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_retrieve
func (c *Client) BGPAutonomousSystemFilter(q *url.Values) ([]AutonomousSystem, error) {
	resp := make([]AutonomousSystem, 0)
	return resp, core.Paginate[AutonomousSystem](c.Client, "plugins/bgp/autonomous-systems/", q, &resp)
}
