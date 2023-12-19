package shared

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/josh-silvas/nbot/shared/nautobot"
)

// DeviceAggregate : Nested device struct, so we can attach methods.
type DeviceAggregate struct {
	Record       nautobot.Device
	ConsolePorts []nautobot.ConsolePort
}

// GetDeviceAggregates : Returns a slice of DeviceAggregate types.
func GetDeviceAggregates(c *nautobot.Client, params url.Values) ([]DeviceAggregate, error) {
	r := make([]DeviceAggregate, 0)
	ds, err := c.GetDevices(&params)
	if err != nil {
		return nil, fmt.Errorf("cfg.Nautobot.Devices:%w", err)
	}
	for i := range ds {
		this := &DeviceAggregate{Record: ds[i], ConsolePorts: make([]nautobot.ConsolePort, 0)}
		if err := this.getConsolePorts(c); err != nil {
			return nil, fmt.Errorf("cfg.Nautobot.getConsolePorts:%w", err)
		}
		r = append(r, *this)
	}
	return r, nil
}

// getConsolePorts func will return the console ports of a device.
func (d *DeviceAggregate) getConsolePorts(c *nautobot.Client) error {
	p, err := c.ConsoleConnections(&url.Values{"device_id": []string{d.Record.ID}})
	if err != nil {
		return fmt.Errorf("cfg.Nautobot.ConsoleConnections:%w", err)
	}
	d.ConsolePorts = p
	return nil
}

// valuePrimaryIP func will return the primary IP address of a device.
func (d *DeviceAggregate) valuePrimaryIP() string {
	if d.Record.PrimaryIP == nil {
		return ""
	}
	return d.Record.PrimaryIP.Address
}

// valueManufacturer func will return the manufacturer of a device.
func (d *DeviceAggregate) valueManufacturer() string {
	if d.Record.DeviceType.Manufacturer == nil {
		return ""
	}
	return d.Record.DeviceType.Manufacturer.Name
}

// valueTenant func will return the tenant of a device.
func (d *DeviceAggregate) valueTenant() string {
	if d.Record.Tenant == nil {
		return ""
	}
	return d.Record.Tenant.Name
}

// valueTags func will return a string of tags from a slice of tags.
func (d *DeviceAggregate) valueTags() string {
	s := make([]string, 0)
	for i := range d.Record.Tags {
		s = append(s, d.Record.Tags[i].Slug)
	}
	return strings.Join(s, ",")
}

// valueConsole func will return the console port of a device.
func (d *DeviceAggregate) valueConsole() string {
	for _, p := range d.ConsolePorts {
		con := p.ConnectedEndpoint
		if con.Port == 0 {
			continue
		}
		return fmt.Sprintf("%s:%d", con.Device.Name, con.Port+3000)
	}
	return "None"
}
