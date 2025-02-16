package sot

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/dmux/internal/keyring"
	"github.com/manifoldco/promptui"
)

var (
	// SupportedSoT : A map of supported SoT backends.
	SupportedSoT = map[string]string{
		"nautobot": "Nautobot",
		"netbox":   "Netbox",
	}

	// ErrorNotImplemented : Error returned when a method is not implemented.
	ErrorNotImplemented = errors.New("not implemented")
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
		ID       string `dmux:"-"`
		AssetTag string `dmux:"-"`
		Hostname string `dmux:"required"`
		Status   string `dmux:"required"`
		IP       string `dmux:"required"`
		Platform string `dmux:"-"`
		Tenant   string `dmux:"-"`
		Location string `dmux:"-"`
		Serial   string `dmux:"-"`
		Comments string `dmux:"-"`
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
	backend, err := settings.KeyFromSection("dmux", "sot-type", selectSoT)
	if err != nil {
		l.Fatalf("cfg.KeyFromSection(%s)", err)
	}

	switch strings.ToLower(backend.String()) {
	case "nautobot":
		nbURL, err := settings.KeyFromSection("nautobot", "url", setSotURL)
		if err != nil {
			return nil, fmt.Errorf("settings.KeyFromSection: %w", err)
		}
		nKey, err := settings.Nautobot()
		if err != nil {
			l.Fatalf("Nautobot(%s)", err)
		}
		return NewNautobot(nKey.Password(), nbURL.String())
	case "netbox":
		nbURL, err := settings.KeyFromSection("netbox", "url", setSotURL)
		if err != nil {
			return nil, fmt.Errorf("settings.KeyFromSection: %w", err)
		}
		nKey, err := settings.Netbox()
		if err != nil {
			l.Fatalf("Netbox(%s)", err)
		}
		return NewNetbox(nKey.Password(), nbURL.String())
	default:
		return nil, fmt.Errorf("unknown backend: %s", backend)
	}
}

func setSotURL() (string, error) {
	p := promptui.Prompt{
		Label: "Enter the URL for your SoT instance",
	}
	return p.Run()
}

func selectSoT() (string, error) {
	prompt := promptui.Select{
		Label: "Select the backend source-of-truth from the following options. Select",
		Items: func() []string {
			r := make([]string, 0)
			for k := range SupportedSoT {
				r = append(r, SupportedSoT[k])
			}
			return r
		}(),
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	for k := range SupportedSoT {
		if result == SupportedSoT[k] {
			fmt.Println(text.FgHiGreen.Sprintf("Using %s as the backend source-of-truth. "+
				"You can always change your backend system by editing the "+
				"DMux configuration file in $HOME/%s/%s settings. 😊",
				SupportedSoT[k], keyring.ConfigPath, keyring.SettingsFile))
			return k, nil
		}
	}
	return "", fmt.Errorf("SoT not found `%s`", result)
}
