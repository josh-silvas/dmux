package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	// NestedCircuit : defines a stub circuit entry in Nautobot
	NestedCircuit struct {
		ID      string `json:"id"`
		CID     string `json:"cid"`
		Display string `json:"display"`
		URL     string `json:"url"`
	}

	// NestedCircuitType : defines a stub circuit type entry in Nautobot
	NestedCircuitType struct {
		Display string `json:"display"`
		ID      string `json:"id"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	}

	// Circuit : defines a circuit entry in Nautobot
	Circuit struct {
		ID           string            `json:"id"`
		Display      string            `json:"display"`
		URL          string            `json:"url"`
		CircuitID    string            `json:"cid"`
		Provider     NestedProvider    `json:"provider"`
		Type         NestedCircuitType `json:"type"`
		Status       LabelValue        `json:"status"`
		Tenant       NestedTenant      `json:"tenant"`
		InstallDate  string            `json:"install_date"`
		CommitRate   int               `json:"commit_rate"`
		Description  string            `json:"description"`
		TerminationA NestedTermination `json:"termination_a"`
		TerminationZ NestedTermination `json:"termination_z"`
		Comments     string            `json:"comments"`
		Created      string            `json:"created"`
		LastUpdated  time.Time         `json:"last_updated"`
		Tags         []NestedTag       `json:"tags"`
		NotesURL     string            `json:"notes_url"`
		CustomFields map[string]any    `json:"custom_fields"`
	}
)

// GetCircuits : method for API endpoint /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) GetCircuits(query *url.Values) ([]Circuit, error) {
	req, err := c.Request(
		http.MethodGet,
		"circuits/circuits/",
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	resp := new(rawListResponse)
	ret := make([]Circuit, 0)

	if err = c.Do(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("error decoding results field: %w", err)
	}
	return ret, err
}
