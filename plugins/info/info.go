package info

import (
	"errors"
	"net"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/josh-silvas/nbot/nlog"
	"github.com/josh-silvas/nbot/shared/sot"
)

const pluginName = "info"

// Plugin type is used as the command and calling function for each plugin
type Plugin struct {
	core.PluginBase
	argD  *string
	argSe *string
}

func (p Plugin) Register(c *core.Parser) core.PluginIfc {
	p.Log = nlog.NewWithGroup(pluginName)
	p.C = c.NewCommand(pluginName, "Gathers device information.")
	p.argD = p.C.StringPositional(&argparse.Options{Required: true, Help: "One of DeviceName, IPAddress, MacAddress"})
	p.argSe = p.C.String("", "serial", &argparse.Options{Help: "Devices by serial number."})
	return p
}

func (p Plugin) CMD() *argparse.Command {
	return p.C
}

func (p Plugin) Func(cfg keyring.Settings) {
	// 1. Check that a device has been passed in
	if strings.TrimSpace(*p.argD) == "" && strings.TrimSpace(*p.argSe) == "" {
		p.Log.Fatalf("nbot info [Name, IP Address, Mac Address] or [DeviceSerial] must be provided!")
	}

	s, err := sot.New(cfg)
	if err != nil {
		p.Log.Fatalf("error initializing sot.New(%s)", err)
	}

	if *p.argSe != "" {
		device, err := s.GetDevice(sot.BySerial, *p.argSe)
		if err != nil {
			if errors.Is(err, sot.ErrorNotImplemented) {
				p.Log.Warn("sot.BySerial not implemented for this backend.")
				return
			}
			p.Log.Errorf("sot.BySerial: %s", err)
		}
		PrintDevices([]sot.Device{device})
		return
	}

	// No partial matches for IP addresses, return the single device if found.
	if ip := net.ParseIP(*p.argD); ip != nil {
		device, err := s.GetDevice(sot.ByIP, ip)
		if err != nil {
			if errors.Is(err, sot.ErrorNotImplemented) {
				p.Log.Warn("sot.ByIP not implemented for this backend.")
				return
			}
			p.Log.Fatalf("sot.ByIP: %s", err)
			return
		}
		PrintDevices([]sot.Device{device})
		return
	}

	// No partial matches for mac-addresses, return the single device if found.
	if mac, err := net.ParseMAC(*p.argD); err == nil {
		device, err := s.GetDevice(sot.ByMac, mac.String())
		if err != nil {
			if errors.Is(err, sot.ErrorNotImplemented) {
				p.Log.Warn("sot.ByMac not implemented for this backend.")
				return
			}
			p.Log.Fatalf("sot.ByMac: %s", err)
			return
		}
		PrintDevices([]sot.Device{device})
		return
	}

	device, err := s.GetDevice(sot.ByName, *p.argD)
	if err != nil {
		p.Log.Fatalf("sot.ByName: %s", err)
	}
	PrintDevices([]sot.Device{device})
}
