package tenancy

import (
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/url"
	"time"
)

type (
	// Tenant : defines a tenant entry in Nautobot
	Tenant struct {
		ID             string             `json:"id"`
		Display        string             `json:"display"`
		URL            string             `json:"url"`
		Name           string             `json:"name"`
		Slug           string             `json:"slug"`
		Group          nested.TenantGroup `json:"group"`
		Description    string             `json:"description"`
		Comments       string             `json:"comments"`
		CircuitCount   int                `json:"circuit_count"`
		DeviceCount    int                `json:"device_count"`
		IpaddressCount int                `json:"ipaddress_count"`
		PrefixCount    int                `json:"prefix_count"`
		RackCount      int                `json:"rack_count"`
		SiteCount      int                `json:"site_count"`
		VMCount        int                `json:"virtualmachine_count"`
		VlanCount      int                `json:"vlan_count"`
		VrfCount       int                `json:"vrf_count"`
		Created        string             `json:"created"`
		LastUpdated    time.Time          `json:"last_updated"`
		Tags           []extras.Tag       `json:"tags"`
		NotesURL       string             `json:"notes_url"`
		CustomFields   map[string]any     `json:"custom_fields"`
	}
)

// TenantFilter : Go function to process requests for the endpoint: /api/tenancy/tenants/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenants_list
func (c *Client) TenantFilter(q *url.Values) ([]Tenant, error) {
	resp := make([]Tenant, 0)
	return resp, core.Paginate[Tenant](c.Client, "tenancy/tenants/", q, &resp)
}
