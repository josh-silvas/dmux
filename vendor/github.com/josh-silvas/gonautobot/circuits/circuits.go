package circuits

import (
	"github.com/josh-silvas/gonautobot/core"
)

// Client : Abstracted base client to use with Nautobot
type Client struct {
	*core.Client
}

// New : Initializes the BGP client.
func New(r *core.Client) *Client {
	return &Client{r}
}
