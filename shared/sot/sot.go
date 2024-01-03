package sot

import (
	"fmt"
	"strings"

	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/sirupsen/logrus"
)

const (
	// ByIP : Used to search for a device by IP address.
	ByIP = "ip"
	// ByID : Used to search for a device by ID.
	ByID = "id"
	// ByName : Used to search for a device by name.
	ByName = "name"
	// BySerial : Used to search for a device by serial number.
	BySerial = "serial"
	// ByMac : Used to search for a device by MAC address.
	ByMac = "mac"
)

type (
	// SoT : Stored memory objects for the SoT client.
	SoT interface {
		// GetDevice : A method used in the SoT interface to fetch a single device from the
		// backend system using a unique identifier.
		// The return value is a single device, or an error.
		GetDevice(method string, value any) (Device, error)
	}

	// Device : A struct used to represent a device in the SoT.
	Device struct {
		ID       string `nbot:"-"`
		AssetTag string `nbot:"-"`
		Hostname string `nbot:"required"`
		Status   string `nbot:"required"`
		IP       string `nbot:"required"`
		Platform string `nbot:"-"`
		Tenant   string `nbot:"-"`
		Location string `nbot:"-"`
		Serial   string `nbot:"-"`
		Comments string `nbot:"-"`
		Console  Console
	}

	// Console : A struct used to represent a console connection to a device in the SoT.
	Console struct {
		ID       string
		Hostname string
		Port     string
	}
)

// New : Function used to create a new SoT client data type.
func New(settings keyring.Settings) (SoT, error) {
	backend, err := settings.KeyFromSection("nbot", "sot-type", "nautobot")
	if err != nil {
		logrus.Fatalf("cfg.KeyFromSection(%s)", err)
	}

	switch strings.ToLower(backend.String()) {
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
