package sot

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/dcim"
	"net"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/dmux/internal/nlog"
	"github.com/josh-silvas/gonautobot"
	"github.com/manifoldco/promptui"
)

var l = nlog.NewWithGroup("sot")

// Nautobot implementation of the SoT interface.
type Nautobot struct {
	nautobot.Client
}

// NewNautobot : Creates a new Nautobot client.
func NewNautobot(token, nbURL string, opts ...nautobot.Option) (*Nautobot, error) {
	// Add on the default options
	opts = append(opts, nautobot.WithToken(token), nautobot.WithEndpoint(nbURL))
	c := nautobot.New(opts...)
	return &Nautobot{Client: *c}, nil
}

// GetDevice : Nautobot: Returns a single device from NautobotV1.
func (n Nautobot) GetDevice(method string, value any) (Device, error) {
	switch strings.ToLower(method) {
	case ByID:
		return n.deviceByNameOrID(value.(string))
	case ByName:
		return n.deviceByNameOrID(value.(string))
	case BySerial:
		return n.deviceBySerial(value.(string))
	case ByIP:
		return n.deviceByIP(value.(net.IP))
	case ByMac:
		return n.deviceByMac(value.(string))
	default:
		return Device{}, ErrorNotImplemented
	}
}

// deviceByNameOrID : Nautobot: Returns a device by name or UUIDv4 ID.
func (n Nautobot) deviceByNameOrID(name string) (Device, error) {
	// 1. First, attempt to convert the argD variable to an integer,
	//    if successful, we can assume this is a NautobotV1 ID.
	params := make(url.Values)

	if _, err := uuid.Parse(name); err == nil {
		params["id"] = []string{name}
	}

	// 2. Second, if there is still not a NautobotV1 ID, but the name variable is
	//    set with something, then we can include this into the name search in NautobotV1.
	if _, ok := params["id"]; !ok && name != "" {
		params["name__ic"] = []string{name}
	}

	return n.getDevice(&params)
}

// deviceBySerial : Nautobot: Returns a device by serial number.
func (n Nautobot) deviceBySerial(serial string) (Device, error) {
	params := make(url.Values)
	params["serial"] = []string{serial}
	return n.getDevice(&params)
}

// deviceByMac : Nautobot: Returns a device by mac address.
func (n Nautobot) deviceByMac(mac string) (Device, error) {
	params := make(url.Values)
	params["mac_address__ic"] = []string{mac}
	return n.getDevice(&params)
}

// deviceByIP : Nautobot: Returns a device by IP address.
func (n Nautobot) deviceByIP(ip net.IP) (Device, error) {
	// 1. Create the device object, we don't want to fail out
	//    in there's a NautobotV1 error, since we already have what we
	//    need to connect.
	dev := Device{
		IP:       ip.String(),
		Hostname: ip.String(),
	}
	// 2. Fetch all the devices from NautobotV1 that match this IP address.
	ips, err := n.Ipam.IPAddressFilter(&url.Values{"address": []string{ip.String()}})
	if err != nil {
		l.Errorf("[COMERR:NautobotV1:IPAddresses::%s]", err)
		return dev, nil
	}

	// 3. For each IP address, query the first match found for a device
	//    within NautobotV1.
	d := make([]dcim.Device, 0)
	for i := range ips {
		if ips[i].AssignedObject == nil {
			continue
		}
		if ips[i].AssignedObject.Device.ID != "" {
			item, err := n.Dcim.DeviceFilter(
				&url.Values{
					"id": []string{ips[i].AssignedObject.Device.ID},
				},
			)
			if err != nil {
				l.Errorf("[COMMERR:NautobotV1:Devices::%s]", err)
				return dev, nil
			}
			d = append(d, item...)
		}
	}
	if len(d) == 0 {
		fmt.Println(text.FgHiYellow.Sprintf("\nUnable to find NautobotV1 device for %s. Using IP only.", ip.String()))
	}

	dev = nbDeviceToSotDevice(&n.Client, d[0])
	dev.IP = ip.String()

	// 4. Finally, return the NautobotV1 device.
	return dev, nil
}

// deviceByName : Nautobot: Returns a device by name.
func (n Nautobot) getDevice(params *url.Values) (Device, error) {
	// 1. Ignore devices in Offline status
	params.Set("depth", "1")

	// 2. Query NautobotV1 with the newly built query parameters.
	d, err := n.Dcim.DeviceFilter(params)
	if err != nil {
		return Device{}, fmt.Errorf("[Devices::%w]", err)
	}

	// 3. If we found more or less than a single device, we need to prompt
	//    for a single device.
	if len(d) > 1 {
		d[0], err = nbDeviceSelect(d)
		if err != nil {
			return Device{}, fmt.Errorf("[Devices::%w]", err)
		}
	}
	if len(d) == 0 {
		return Device{}, fmt.Errorf("unable to find Nautobot device for `%v`", params)
	}
	switch {
	case d[0].PrimaryIP != nil:
		break
	case d[0].PrimaryIP4 != nil:
		d[0].PrimaryIP = d[0].PrimaryIP4
	case d[0].PrimaryIP6 != nil:
		d[0].PrimaryIP = d[0].PrimaryIP6
	default:
		return Device{}, fmt.Errorf("`%v` does not have a primary IP assigned in Nautobot", params)
	}

	// 5. Finally, return the device.
	return nbDeviceToSotDevice(&n.Client, d[0]), nil
}

func nbConsolePort(n *nautobot.Client, id string) (dcim.ConsolePort, error) {
	p, err := n.Dcim.ConsoleConnectionFilter(&url.Values{"device_id": []string{id}})
	if err != nil {
		return dcim.ConsolePort{}, fmt.Errorf("cfg.NautobotV1.ConsoleConnections:%w", err)
	}
	if len(p) == 0 {
		return dcim.ConsolePort{}, fmt.Errorf("no console ports found for device `%s`", id)
	}
	return p[0], nil
}

func nbDeviceToSotDevice(n *nautobot.Client, d dcim.Device) Device {
	ret := Device{
		Hostname: d.Name,
		IP: func() string {
			a := d.PrimaryIP.Address
			if strings.Contains(a, "/") {
				a = strings.Split(a, "/")[0]
			}
			return a
		}(),
		Platform: d.DeviceType.Model,
		Location: func() string {
			if d.Site != nil {
				return d.Site.Name
			}
			if d.Location != nil {
				return d.Location.Name
			}
			return ""
		}(),
		Status:   d.Status.Value,
		Serial:   d.Serial,
		AssetTag: d.AssetTag,
		Tenant:   d.Tenant.Name,
	}
	if p, err := nbConsolePort(n, d.ID); err == nil {
		ret.Console = Console{
			Hostname: p.ConnectedEndpoint.Device.Name,
			Port:     p.ConnectedEndpoint.Name,
			ID:       p.Device.ID,
		}
	}
	return ret
}

func nbDeviceSelect(devices []dcim.Device) (dcim.Device, error) {
	prompt := promptui.Select{
		Label: "Multiple devices found. Select",
		Items: func() []string {
			r := make([]string, 0)
			for k := range devices {
				r = append(r, devices[k].Name)
			}
			return r
		}(),
	}

	_, result, err := prompt.Run()
	if err != nil {
		return dcim.Device{}, err
	}
	for k := range devices {
		if result == devices[k].Name {
			return devices[k], nil
		}
	}
	return dcim.Device{}, fmt.Errorf("unable to determine device from `%s`", result)
}
