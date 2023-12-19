package nautobot

type (
	// NestedCable : defines a stub cable entry in Nautobot
	NestedCable struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Label   string `json:"label"`
		URL     string `json:"url"`
	}
)
