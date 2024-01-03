package nautobot_v1

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-cleanhttp"

	"github.com/sirupsen/logrus"
)

type (
	options struct {
		baseURL       *string
		httpTimeout   *int
		logLevel      *logrus.Level
		proxyURL      *string
		useDev        *bool
		httpTransport *http.Transport
	}

	// Option : Basic options allowed with this client.
	Option func(*options)
)

// SetDevEnv : This option can force client to use the dev env.
func SetDevEnv(b bool) Option {
	return func(o *options) {
		o.useDev = &b
	}
}

// SetHTTPTimeout : This option will set a custom http timeout
func SetHTTPTimeout(timeout int) Option {
	return func(o *options) {
		o.httpTimeout = &timeout
	}
}

// SetHTTPTransport : This option will set a custom http transport
func SetHTTPTransport(transport *http.Transport) Option {
	return func(o *options) {
		o.httpTransport = transport
	}
}

// WithLogLevel : set custom client logging level
func WithLogLevel(lvl logrus.Level) Option {
	return func(o *options) {
		o.logLevel = &lvl
	}
}

// WithProxy : Will pass in a proxy URL to the init function.
func WithProxy(url string) Option {
	return func(o *options) {
		o.proxyURL = &url
	}
}

// WithURL : Will pass in a base URL to the init function.
func WithURL(baseURL string) Option {
	return func(o *options) {
		o.baseURL = &baseURL
	}
}

func (c *Client) processOptions(opts ...Option) error {
	options := new(options)
	for _, opt := range opts {
		opt(options)
	}

	c.client = cleanhttp.DefaultPooledClient()

	if options.httpTimeout != nil {
		c.client.Timeout = time.Duration(*options.httpTimeout) * time.Second
	}

	if options.httpTransport != nil {
		c.client.Transport = options.httpTransport
	}

	if options.proxyURL != nil {
		pURL, err := url.Parse(*options.proxyURL)
		if err != nil {
			logrus.Fatal(err)
		}
		t := cleanhttp.DefaultPooledTransport()
		t.Proxy = http.ProxyURL(pURL)
		c.client.Transport = t
	}

	// WithURL() overrides UseDev()
	if options.baseURL != nil {
		var err error
		c.instance, err = url.Parse(*options.baseURL)
		if err != nil {
			return fmt.Errorf("error parsing baseURL: %w", err)
		}
	}

	c.log = func() *logrus.Logger {
		logger := logrus.New()
		logger.Level = logrus.InfoLevel
		if options.logLevel != nil {
			logger.Level = *options.logLevel
		}
		return logger
	}()

	return nil
}
