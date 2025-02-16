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
		Family             *shared.LabelValueInt    `json:"family"`
		LastUpdated        string                   `json:"last_updated"`
		NATInside          *string                  `json:"nat_inside"`
		NATOutside         *string                  `json:"nat_outside"`
		NotesURL           string                   `json:"notes_url"`
		Role               *shared.LabelValue       `json:"role"`
		Status             *shared.LabelValue       `json:"status"`
		Tags               []extras.Tag             `json:"tags"`
		Tenant             *nested.Tenant           `json:"tenant"`
		URL                string                   `json:"url"`
		VRF                *nested.VRF              `json:"vrf"`
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

// IPAddressGet : Go function to process requests for the endpoint: /api/ipam/ip_addresses/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_ip_addresses_retrieve
func (c *Client) IPAddressGet(uuid string) (*IPAddress, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/ip-addresses/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(IPAddress)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// IPAddressFilter : Go function to process requests for the endpoint: /api/ipam/ip_addresses/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_ip_addresses_list
func (c *Client) IPAddressFilter(q *url.Values) ([]IPAddress, error) {
	resp := make([]IPAddress, 0)
	return resp, core.Paginate[IPAddress](c.Client, "ipam/ip-addresses/", q, &resp)
}
