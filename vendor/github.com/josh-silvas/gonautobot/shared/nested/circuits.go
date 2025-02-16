package nested

type (
	// Provider : Models a subset on the Provider model for nested responses.
	Provider struct {
		CircuitCount int    `json:"circuit_count"`
		Display      string `json:"display"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Slug         string `json:"slug"`
		URL          string `json:"url"`
	}

	// CircuitType : Models a subset on the CircuitType model for nested responses.
	CircuitType struct {
		Display string `json:"display"`
		ID      string `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	}
)
