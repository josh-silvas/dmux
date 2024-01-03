package nautobotv1

// VRF : defines a stub vrf entry in Nautobot
type VRF struct {
	ID          string  `json:"id"`
	Display     string  `json:"display"`
	Name        string  `json:"name"`
	PrefixCount int     `json:"prefix_count"`
	RD          *string `json:"rd"`
	URL         string  `json:"url"`
}
