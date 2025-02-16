package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// Interface : defines an interface as represented in Nautobot
	//
	// CablePeer and ConnectedEndpoint must be decoded dynamically; they are dependent
	// on the value of 'CablePeerType' and 'ConnectedEndpointType', e.g., "dcim.interface"
	Interface struct {
		ID                         string                 `json:"id"`
		Bridge                     *Interface             `json:"bridge"`
		Cable                      *nested.Cable          `json:"cable"`
		CablePeer                  json.RawMessage        `json:"cable_peer"`
		CablePeerType              *string                `json:"cable_peer_type"`
		ConnectedEndpoint          json.RawMessage        `json:"connected_endpoint"`
		ConnectedEndpointReachable bool                   `json:"connected_endpoint_reachable"`
		ConnectedEndpointType      *string                `json:"connected_endpoint_type"`
		CountIPAddresses           int                    `json:"count_ipaddresses"`
		Created                    string                 `json:"created"`
		CustomFields               map[string]interface{} `json:"custom_fields"`
		Description                string                 `json:"description"`
		Device                     *Device                `json:"device"`
		Display                    string                 `json:"display"`
		Enabled                    bool                   `json:"enabled"`
		Label                      string                 `json:"label"`
		Lag                        *Interface             `json:"lag"`
		LastUpdated                string                 `json:"last_updated"`
		MACAddress                 *string                `json:"mac_address"`
		MgmtOnly                   bool                   `json:"mgmt_only"`
		Mode                       *shared.LabelValue     `json:"mode"`
		MTU                        *int                   `json:"mtu"`
		Name                       string                 `json:"name"`
		NotesURL                   string                 `json:"notes_url"`
		ParentInterface            *Interface             `json:"parent_interface"`
		TaggedVLANs                []nested.VLAN          `json:"tagged_vlans"`
		Tags                       []extras.Tag           `json:"tags"`
		Type                       *shared.LabelValue     `json:"type"`
		UntaggedVLAN               *nested.VLAN           `json:"untagged_vlan"`
		URL                        string                 `json:"url"`
	}
)

// InterfaceGet : Go function to process requests for the endpoint: /api/dcim/interfaces/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_retrieve
func (c *Client) InterfaceGet(uuid string) (*Interface, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/interfaces/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(Interface)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// InterfaceFilter : Go function to process requests for the endpoint: /api/dcim/interfaces/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_list
func (c *Client) InterfaceFilter(q *url.Values) ([]Interface, error) {
	resp := make([]Interface, 0)
	if err := core.Paginate[Interface](c.Client, "dcim/interfaces/", q, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
