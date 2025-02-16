package tenancy

import (
	"github.com/josh-silvas/gonautobot/core"
	"net/url"
	"time"
)

type (
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

// TenantGroupFilter : Go function to process requests for the endpoint: /api/tenancy/tenant-groups/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenant_groups_list
func (c *Client) TenantGroupFilter(q *url.Values) ([]TenantGroup, error) {
	resp := make([]TenantGroup, 0)
	return resp, core.Paginate[TenantGroup](c.Client, "tenancy/tenant-groups/", q, &resp)
}
