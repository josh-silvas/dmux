package nautobot

import "time"

type (
	// NestedCircuitTermination : defines a stub circuit-termination entry in Nautobot
	NestedCircuitTermination struct {
		ID       string         `json:"id"`
		Cable    string         `json:"cable"`
		Circuit  *NestedCircuit `json:"circuit"`
		Display  string         `json:"display"`
		TermSide string         `json:"term_side"`
		URL      string         `json:"url"`
	}

	// NestedTermination : defines a stub termination entry in Nautobot
	NestedTermination struct {
		Display           string      `json:"display"`
		ID                string      `json:"id"`
		URL               string      `json:"url"`
		Site              NestedSite  `json:"site"`
		ProviderNetwork   interface{} `json:"provider_network"`
		ConnectedEndpoint struct {
			Display string       `json:"display"`
			ID      string       `json:"id"`
			URL     string       `json:"url"`
			Device  NestedDevice `json:"device"`
			Name    string       `json:"name"`
			Cable   string       `json:"cable"`
		} `json:"connected_endpoint"`
		PortSpeed     int       `json:"port_speed"`
		UpstreamSpeed int       `json:"upstream_speed"`
		XConnectID    string    `json:"xconnect_id"`
		Created       string    `json:"created"`
		LastUpdated   time.Time `json:"last_updated"`
		NotesURL      string    `json:"notes_url"`
	}
)
