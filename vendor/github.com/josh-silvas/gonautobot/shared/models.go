package shared

import (
	"encoding/json"
)

type (
	// LabelValue defines a repeating structure used throughout the API
	LabelValue struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}

	// LabelValueInt defines a repeating structure used throughout the API
	LabelValueInt struct {
		Label string `json:"label"`
		Value int    `json:"value"`
	}

	// ResponseList is an abstracted response when a list is returned in Nautobot.
	ResponseList struct {
		Count    int             `json:"count"`
		Next     string          `json:"next"`
		Previous string          `json:"previous"`
		Results  json.RawMessage `json:"results"`
	}

	// BulkDelete is used to for bulk-delete operations in Nautobot.
	BulkDelete []struct {
		ID string `json:"id"`
	}
)
