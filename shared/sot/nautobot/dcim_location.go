package nautobot

type (
	// NestedLocation : defines a stub location entry in Nautobot
	NestedLocation struct {
		ID        string `json:"id"`
		Display   string `json:"display"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		TreeDepth int    `json:"tree_depth"`
		URL       string `json:"url"`
	}
)
