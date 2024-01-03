package nautobot_v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	// Tenant : defines a tenant entry in Nautobot
	Tenant struct {
		ID             string         `json:"id"`
		Display        string         `json:"display"`
		URL            string         `json:"url"`
		Name           string         `json:"name"`
		Slug           string         `json:"slug"`
		Group          TenantGroup    `json:"group"`
		Description    string         `json:"description"`
		Comments       string         `json:"comments"`
		CircuitCount   int            `json:"circuit_count"`
		DeviceCount    int            `json:"device_count"`
		IpaddressCount int            `json:"ipaddress_count"`
		PrefixCount    int            `json:"prefix_count"`
		RackCount      int            `json:"rack_count"`
		SiteCount      int            `json:"site_count"`
		VMCount        int            `json:"virtualmachine_count"`
		VlanCount      int            `json:"vlan_count"`
		VrfCount       int            `json:"vrf_count"`
		Created        string         `json:"created"`
		LastUpdated    time.Time      `json:"last_updated"`
		Tags           []Tag          `json:"tags"`
		NotesURL       string         `json:"notes_url"`
		CustomFields   map[string]any `json:"custom_fields"`
	}
	// TenantGroup : defines a tenant-group entry in Nautobot
	TenantGroup struct {
		ID           string         `json:"id"`
		Display      string         `json:"display"`
		URL          string         `json:"url"`
		Name         string         `json:"name"`
		Slug         string         `json:"slug"`
		Parent       *TenantGroup   `json:"parent"`
		Description  string         `json:"description"`
		TenantCount  int            `json:"tenant_count"`
		Depth        int            `json:"_depth"`
		Created      string         `json:"created"`
		LastUpdated  time.Time      `json:"last_updated"`
		NotesURL     string         `json:"notes_url"`
		CustomFields map[string]any `json:"custom_fields"`
	}
)

// GetTenants : method for API endpoint /api/tenancy/tenants/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenants_list
func (c *Client) GetTenants(query *url.Values) ([]Tenant, error) {
	req, err := c.Request(
		http.MethodGet,
		"tenancy/tenants/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Tenant, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
