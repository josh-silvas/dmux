package nautobot

// Rack : Type definition for a Nautobot Rack
type Rack struct {
	ID          string `json:"id"`
	DeviceCount int    `json:"device_count"`
	Display     string `json:"display"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}
