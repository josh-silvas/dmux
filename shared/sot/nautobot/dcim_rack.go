package nautobot

type (
	// NestedRack : defines a stub rack entry in Nautobot
	NestedRack struct {
		ID          string `json:"id"`
		DeviceCount int    `json:"device_count"`
		Display     string `json:"display"`
		Name        string `json:"name"`
		URL         string `json:"url"`
	}
)
