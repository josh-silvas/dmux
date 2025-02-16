package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// Site : Data representation of a Site model in Nautobot.
	Site struct {
		ASN                 *int               `json:"asn"`
		CircuitCount        int                `json:"circuit_count"`
		Comments            string             `json:"comments" datastore:",noindex"`
		ContactEmail        string             `json:"contact_email"`
		ContactName         string             `json:"contact_name"`
		ContactPhone        string             `json:"contact_phone"`
		Created             string             `json:"created"`
		CustomFields        map[string]any     `json:"custom_fields"`
		Description         string             `json:"description"`
		DeviceCount         int                `json:"device_count"`
		Display             string             `json:"display"`
		Facility            string             `json:"facility"`
		ID                  string             `json:"id"`
		LastUpdated         string             `json:"last_updated"`
		Latitude            json.RawMessage    `json:"latitude"`
		Longitude           json.RawMessage    `json:"longitude"`
		Name                string             `json:"name"`
		NotesURL            string             `json:"notes_url"`
		PhysicalAddress     string             `json:"physical_address"`
		PrefixCount         int                `json:"prefix_count"`
		RackCount           int                `json:"rack_count"`
		Region              *Region            `json:"region"`
		ShippingAddress     string             `json:"shipping_address"`
		Slug                string             `json:"slug"`
		Status              *shared.LabelValue `json:"status"`
		Tags                []extras.Tag       `json:"tags"`
		Tenant              *nested.Tenant     `json:"tenant"`
		TimeZone            json.RawMessage    `json:"time_zone"`
		URL                 string             `json:"url"`
		VirtualMachineCount int                `json:"virtualmachine_count"`
		VLANCount           int                `json:"vlan_count"`
	}
)

// SiteGet : Go function to process requests for the endpoint: /api/dcim/sites/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_sites_retrieve
func (c *Client) SiteGet(uuid string) (*Site, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/sites/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(Site)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// SiteFilter : Go function to process requests for the endpoint: /api/dcim/sites/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_sites_list
func (c *Client) SiteFilter(q *url.Values) ([]Site, error) {
	req, err := c.Request(http.MethodGet, "dcim/sites/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]Site, 0)

	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("GetSites.error.json.Unmarshal(%w)", err)
	}
	return ret, err
}
