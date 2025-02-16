package nested

type (

	// ConfigContextSchema : Models a subset on the ConfigContextSchema model for nested responses.
	ConfigContextSchema struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}

	// SecretsGroup : Models a subset on the SecretsGroup model for nested responses.
	SecretsGroup struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}
)
