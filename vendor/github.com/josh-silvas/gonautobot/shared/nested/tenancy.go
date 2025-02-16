package nested

type (
	// TenantGroup : Models a subset on the TenantGroup model for nested responses.
	TenantGroup struct {
		Display string `json:"display"`
		ID      string `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		Depth   int    `json:"_depth"`
	}

	// Tenant : Models a subset on the Tenant model for nested responses.
	Tenant struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}
)
