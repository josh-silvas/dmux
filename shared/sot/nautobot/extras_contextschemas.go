package nautobot

type (
	// NestedConfigContextSchema : defines a stub context_schema entry in Nautobot
	NestedConfigContextSchema struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}
)
