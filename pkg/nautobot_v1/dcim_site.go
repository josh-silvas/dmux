package nautobot_v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Site : Type represents a site object in Nautobot
type Site struct {
	ASN                 *int            `json:"asn"`
	CircuitCount        int             `json:"circuit_count"`
	Comments            string          `json:"comments" datastore:",noindex"`
	ContactEmail        string          `json:"contact_email"`
	ContactName         string          `json:"contact_name"`
	ContactPhone        string          `json:"contact_phone"`
	Created             string          `json:"created"`
	CustomFields        map[string]any  `json:"custom_fields"`
	Description         string          `json:"description"`
	DeviceCount         int             `json:"device_count"`
	Display             string          `json:"display"`
	Facility            string          `json:"facility"`
	ID                  string          `json:"id"`
	LastUpdated         string          `json:"last_updated"`
	Latitude            json.RawMessage `json:"latitude"`
	Longitude           json.RawMessage `json:"longitude"`
	Name                string          `json:"name"`
	NotesURL            string          `json:"notes_url"`
	PhysicalAddress     string          `json:"physical_address"`
	PrefixCount         int             `json:"prefix_count"`
	RackCount           int             `json:"rack_count"`
	Region              *Region         `json:"region"`
	ShippingAddress     string          `json:"shipping_address"`
	Slug                string          `json:"slug"`
	Status              *LabelValue     `json:"status"`
	Tags                []Tag           `json:"tags"`
	Tenant              *Tenant         `json:"tenant"`
	TimeZone            json.RawMessage `json:"time_zone"`
	URL                 string          `json:"url"`
	VirtualMachineCount int             `json:"virtualmachine_count"`
	VLANCount           int             `json:"vlan_count"`
}

// GetSite : method for API endpoint /api/dcim/sites/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_sites_retrieve
func (c *Client) GetSite(uuid string, query *url.Values) (*Site, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("dcim/sites/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(Site)
	err = c.Do(req, ret)
	return ret, err
}

// GetSites : method for API endpoint /api/dcim/sites/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_sites_list
func (c *Client) GetSites(query *url.Values) ([]Site, error) {
	req, err := c.Request(
		http.MethodGet,
		"dcim/sites/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Site, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
