package nautobot

type (
	// NestedVRF : defines a stub vrf entry in Nautobot
	NestedVRF struct {
		ID          string  `json:"id"`
		Display     string  `json:"display"`
		Name        string  `json:"name"`
		PrefixCount int     `json:"prefix_count"`
		RD          *string `json:"rd"` // Route Distinguisher
		URL         string  `json:"url"`
	}
)
