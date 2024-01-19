package nautobot

// Region : Type represents a region object in Nautobot
type Region struct {
	ID           string         `json:"id"`
	Created      string         `json:"created"`
	CustomFields map[string]any `json:"custom_fields"`
	Depth        int            `json:"_depth"`
	Description  string         `json:"description"`
	Display      string         `json:"display"`
	LastUpdated  string         `json:"last_updated"`
	Name         string         `json:"name"`
	NotesURL     string         `json:"notes_url"`
	Parent       *Region        `json:"parent"`
	SiteCount    int            `json:"site_count"`
	Slug         string         `json:"slug"`
	URL          string         `json:"url"`
}
