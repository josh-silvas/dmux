package nautobotv1

// Tag : defines a tag entry in Nautobot
type Tag struct {
	ID          string `json:"id"`
	Color       string `json:"color"`
	Created     string `json:"created"`
	Display     string `json:"display"`
	LastUpdated string `json:"last_updated"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         string `json:"url"`
}
