package circuits

import (
	"github.com/josh-silvas/gonautobot/shared/nested"
	"time"
)

type (

	// Termination : Models the Termination Endpoint in Nautobot.
	Termination struct {
		Display           string      `json:"display"`
		ID                string      `json:"id"`
		URL               string      `json:"url"`
		Site              nested.Site `json:"site"`
		ProviderNetwork   interface{} `json:"provider_network"`
		ConnectedEndpoint struct {
			Display string        `json:"display"`
			ID      string        `json:"id"`
			URL     string        `json:"url"`
			Device  nested.Device `json:"device"`
			Name    string        `json:"name"`
			Cable   string        `json:"cable"`
		} `json:"connected_endpoint"`
		PortSpeed     int       `json:"port_speed"`
		UpstreamSpeed int       `json:"upstream_speed"`
		XConnectID    string    `json:"xconnect_id"`
		Created       string    `json:"created"`
		LastUpdated   time.Time `json:"last_updated"`
		NotesURL      string    `json:"notes_url"`
	}
)
