package nautobot

import (
	"encoding/json"
	"fmt"
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
		Bridge                     *NestedInterface       `json:"bridge"`
		Cable                      *NestedCable           `json:"cable"`
		CablePeer                  json.RawMessage        `json:"cable_peer"`
		CablePeerType              *string                `json:"cable_peer_type"`
		ConnectedEndpoint          json.RawMessage        `json:"connected_endpoint"`
		ConnectedEndpointReachable bool                   `json:"connected_endpoint_reachable"`
		ConnectedEndpointType      *string                `json:"connected_endpoint_type"`
		CountIPAddresses           int                    `json:"count_ipaddresses"`
		Created                    string                 `json:"created"`
		CustomFields               map[string]interface{} `json:"custom_fields"`
		Description                string                 `json:"description"`
		Device                     *NestedDevice          `json:"device"`
		Display                    string                 `json:"display"`
		Enabled                    bool                   `json:"enabled"`
		Label                      string                 `json:"label"`
		Lag                        *NestedInterface       `json:"lag"`
		LastUpdated                string                 `json:"last_updated"`
		MACAddress                 *string                `json:"mac_address"`
		MgmtOnly                   bool                   `json:"mgmt_only"`
		Mode                       *LabelValue            `json:"mode"`
		MTU                        *int                   `json:"mtu"`
		Name                       string                 `json:"name"`
		NotesURL                   string                 `json:"notes_url"`
		ParentInterface            *NestedInterface       `json:"parent_interface"`
		TaggedVLANs                []NestedVLAN           `json:"tagged_vlans"`
		Tags                       []NestedTag            `json:"tags"`
		Type                       *LabelValue            `json:"type"`
		UntaggedVLAN               *NestedVLAN            `json:"untagged_vlan"`
		URL                        string                 `json:"url"`
	}

	// NestedInterface : defines a stub interface entry in Nautobot
	NestedInterface struct {
		ID      string        `json:"id"`
		Cable   string        `json:"cable"`
		Device  *NestedDevice `json:"device"`
		Display string        `json:"display"`
		Name    string        `json:"name"`
		URL     string        `json:"url"`
	}
)

// GetInterface : method for API endpoint /api/dcim/interfaces/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_retrieve
func (c *Client) GetInterface(uuid string, query *url.Values) (*Interface, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("dcim/interfaces/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(Interface)
	err = c.Do(req, ret)
	return ret, err
}

// GetInterfaces : method for API endpoint /api/dcim/interfaces/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_list
func (c *Client) GetInterfaces(query *url.Values) ([]Interface, error) {
	req, err := c.Request(
		http.MethodGet,
		"dcim/interfaces/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Interface, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
