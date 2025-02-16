package nested

type (

	// Prefix : Models a subset on the Prefix model for nested responses.
	Prefix struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	// IPAddress : Models a subset on the IPAddress model for nested responses.
	IPAddress struct {
		ID      string `json:"id"`
		Address string `json:"address"`
		Display string `json:"display"`
		Family  int    `json:"family"`
		URL     string `json:"url"`
	}

	// VLAN : Models a subset on the VLAN model for nested responses.
	VLAN struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		VID     int    `json:"vid"`
	}

	// Role : Models a subset on the Role model for nested responses.
	Role struct {
		ID          string `json:"id"`
		Display     string `json:"display"`
		Model       string `json:"model"`
		PrefixCount int    `json:"prefix_count"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
		VLANCount   int    `json:"vlan_count"`
	}

	// VRF : Models a subset on the VRF model for nested responses.
	VRF struct {
		ID          string  `json:"id"`
		Display     string  `json:"display"`
		Name        string  `json:"name"`
		PrefixCount int     `json:"prefix_count"`
		RD          *string `json:"rd"`
		URL         string  `json:"url"`
	}
)
