package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	// IPAddress : defines an IP Address entry in Nautobot
	//
	// AssignedObject will need to be decoded dynamically based
	// on the 'assigned_object_type', e.g., "dcim.interface"
	IPAddress struct {
		ID                 string                   `json:"id"`
		Address            string                   `json:"address"`
		AssignedObject     *AssignedObjectInterface `json:"assigned_object"`
		AssignedObjectID   *string                  `json:"assigned_object_id"`
		AssignedObjectType *string                  `json:"assigned_object_type"`
		Created            string                   `json:"created"`
		CustomFields       map[string]interface{}   `json:"custom_fields"`
		Description        string                   `json:"description"`
		Display            string                   `json:"display"`
		DNSName            string                   `json:"dns_name"`
		Family             *LabelValueInt           `json:"family"`
		LastUpdated        string                   `json:"last_updated"`
		NATInside          *string                  `json:"nat_inside"`
		NATOutside         *string                  `json:"nat_outside"`
		NotesURL           string                   `json:"notes_url"`
		Role               *LabelValue              `json:"role"`
		Status             *LabelValue              `json:"status"`
		Tags               []NestedTag              `json:"tags"`
		Tenant             *NestedTenant            `json:"tenant"`
		URL                string                   `json:"url"`
		VRF                *NestedVRF               `json:"vrf"`
	}

	// NestedIPAddress : defines a stub ipAddress entry in Nautobot
	NestedIPAddress struct {
		ID      string `json:"id"`
		Address string `json:"address"`
		Display string `json:"display"`
		Family  int    `json:"family"`
		URL     string `json:"url"`
	}

	// AssignedObjectInterface : struct type for the `dcim.interface` and `virtualization.vminterface`
	// assigned_object_type in the IPAddress struct.
	// See below for available types:
	// https://github.com/nautobot/nautobot/blob/v1.5.16/nautobot/ipam/constants.py#L35
	AssignedObjectInterface struct {
		Display string `json:"display"`
		ID      string `json:"id"`
		URL     string `json:"url"`
		Device  struct {
			Display string `json:"display"`
			ID      string `json:"id"`
			URL     string `json:"url"`
			Name    string `json:"name"`
		} `json:"device,omitempty"`
		VirtualMachine struct {
			Display string `json:"display"`
			ID      string `json:"id"`
			URL     string `json:"url"`
			Name    string `json:"name"`
		} `json:"virtual_machine,omitempty"`
		Name  string `json:"name"`
		Cable string `json:"cable"`
	}
)

// GetIPAddress : method for API endpoint /api/ipam/ip_addresses/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_ip_addresses_retrieve
func (c *Client) GetIPAddress(uuid string, query *url.Values) (*IPAddress, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("ipam/ip-addresses/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(IPAddress)
	err = c.Do(req, ret)
	return ret, err
}

// GetIPAddresses : method for API endpoint /api/ipam/ip_addresses/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_ip_addresses_list
func (c *Client) GetIPAddresses(query *url.Values) ([]IPAddress, error) {
	req, err := c.Request(
		http.MethodGet,
		"ipam/ip-addresses/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]IPAddress, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
