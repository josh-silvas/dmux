package nautobot

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/circuits"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/dcim"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/ipam"
	"github.com/josh-silvas/gonautobot/plugins"
	"github.com/josh-silvas/gonautobot/tenancy"
	"github.com/josh-silvas/gonautobot/virtualization"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strings"
)

const (
	envNautobotURL   = "NAUTOBOT_URL"
	envNautobotToken = "NAUTOBOT_TOKEN"
)

type (
	// Client : Stored memory objects for the Nautobot client.
	Client struct {
		Request *core.Client

		Circuits       *circuits.Client
		Dcim           *dcim.Client
		Extras         *extras.Client
		Ipam           *ipam.Client
		Plugins        *plugins.Client
		Tenancy        *tenancy.Client
		Virtualization *virtualization.Client
	}

	// Option : Basic options allowed with this client.
	Option func(*Client)
)

// WithEndpoint : Sets the API endpoint to use in Nautobot.
// Default is from NAUTOBOT_URL environment or https://demo.nautobot.com is not specified.
func WithEndpoint(url string) Option {
	return func(o *Client) {
		o.Request.URL = SanitizeURL(url)
	}
}

// WithToken : Sets the token for Nautobot.
// Default is from NAUTOBOT_TOKEN, will error if not either passed in or set as an environment variable.
func WithToken(token string) Option {
	return func(o *Client) {
		o.Request.Token = token
	}
}

// WithHTTPClient : Overrides the default HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(o *Client) {
		o.Request.Client = client
	}
}

// New : Function used to create a new Nautobot client data type.
func New(opts ...Option) *Client {
	c := &Client{
		Request: &core.Client{
			Token: os.Getenv(envNautobotToken),
			URL: func() string {
				r := os.Getenv(envNautobotURL)
				if r == "" {
					r = "https://demo.nautobot.com/api/"
				}
				return r
			}(),
			Client: http.DefaultClient,
		},
	}
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if strings.EqualFold(os.Getenv("DEBUG"), "true") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.Request.Token == "" {
		log.Error().Msg(fmt.Sprintf("Token must be set with core.WithToken('aaaa') or `%s` environment variable", envNautobotToken))
		os.Exit(1)
	}

	c.Circuits = circuits.New(c.Request)
	c.Dcim = dcim.New(c.Request)
	c.Extras = extras.New(c.Request)
	c.Ipam = ipam.New(c.Request)
	c.Plugins = plugins.New(c.Request)
	c.Tenancy = tenancy.New(c.Request)
	c.Virtualization = virtualization.New(c.Request)
	return c
}

// SanitizeURL : Sanitizes a URL for use with Nautobot.
func SanitizeURL(s string) string {
	if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		s = fmt.Sprintf("https://%s", s)
	}
	if !strings.HasSuffix(s, "/") {
		s = fmt.Sprintf("%s/", s)
	}
	if !strings.HasSuffix(s, "api/") {
		s = fmt.Sprintf("%sapi/", s)
	}
	return s
}
