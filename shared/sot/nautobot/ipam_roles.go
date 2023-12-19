package nautobot

type (
	// NestedIPAMRole : defines a stub ipam_role entry in Nautobot
	NestedIPAMRole struct {
		ID          string `json:"id"`
		Display     string `json:"display"`
		Model       string `json:"model"`
		PrefixCount int    `json:"prefix_count"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
		VLANCount   int    `json:"vlan_count"`
	}
)
