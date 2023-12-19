package sot

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/sirupsen/logrus"
)

type (
	SoT interface {
		Devices(values *url.Values) ([]Device, error)
		DeviceByName(name string) (Device, error)
		DeviceByIP(ip net.IP) (Device, error)
	}

	Device struct {
		ID       string
		Hostname string
		IP       string
		Platform string
		Location string
		Comments string
	}
)

func New(backend string, settings keyring.Settings) (SoT, error) {
	switch strings.ToLower(backend) {
	case "nautobot":
		nbURL, err := settings.KeyFromSection("nautobot", "url", "https://demo.nautobot.com")
		if err != nil {
			return nil, fmt.Errorf("settings.KeyFromSection: %s", err)
		}
		nKey, err := settings.Nautobot()
		if err != nil {
			logrus.Fatalf("Nautobot(%s)", err)
		}
		return NewNautobot(nKey.Password(), nbURL.String())
	default:
		return nil, fmt.Errorf("unknown backend: %s", backend)
	}
}
