package plugins

import (
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/plugins/bgp"
)

// Client : Abstracted base client to use with Nautobot
type Client struct {
	BGP *bgp.Client
}

// New : Initializes the BGP client.
func New(c *core.Client) *Client {
	return &Client{
		BGP: bgp.New(c),
	}
}
