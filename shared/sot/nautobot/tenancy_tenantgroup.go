package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	// NestedTenantGroup : defines a stub tenant-group entry in Nautobot
	NestedTenantGroup struct {
		Display string `json:"display"`
		ID      string `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		Depth   int    `json:"_depth"`
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

// GetTenantGroups : method for API endpoint /api/tenancy/tenant-groups/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenant_groups_list
func (c *Client) GetTenantGroups(query *url.Values) ([]TenantGroup, error) {
	req, err := c.Request(
		http.MethodGet,
		"tenancy/tenant-groups/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]TenantGroup, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
