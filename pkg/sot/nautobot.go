package sot

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/pkg/nautobotv1"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

// Nautobot implementation of the SoT interface.
type Nautobot struct {
	nautobotv1.Client
}

// NewNautobot : Creates a new Nautobot client.
func NewNautobot(token, nbURL string, opts ...nautobotv1.Option) (*Nautobot, error) {
	c, err := nautobotv1.New(token, nbURL, opts...)
	if err != nil {
		return nil, err
	}
	return &Nautobot{Client: *c}, nil
}

// GetDevice : Nautobot: Returns a single device from Nautobot.
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
	//    if successful, we can assume this is a Nautobot ID.
	params := make(url.Values)

	if _, err := uuid.Parse(name); err == nil {
		params["id"] = []string{name}
	}

	// 2. Second, if there is still not a Nautobot ID, but the name variable is
	//    set with something, then we can include this into the name search in Nautobot.
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
	//    in there's a Nautobot error, since we already have what we
	//    need to connect.
	dev := Device{
		IP:       ip.String(),
		Hostname: ip.String(),
	}
	// 2. Fetch all the devices from Nautobot that match this IP address.
	ips, err := n.GetIPAddresses(&url.Values{"address": []string{ip.String()}})
	if err != nil {
		logrus.Errorf("[COMERR:Nautobot:IPAddresses::%s]", err)
		return dev, nil
	}

	// 3. For each IP address, query the first match found for a device
	//    within Nautobot.
	d := make([]nautobotv1.Device, 0)
	for i := range ips {
		if ips[i].AssignedObject == nil {
			continue
		}
		if ips[i].AssignedObject.Device.ID != "" {
			item, err := n.GetDevices(
				&url.Values{
					"id": []string{ips[i].AssignedObject.Device.ID},
				},
			)
			if err != nil {
				logrus.Errorf("[COMMERR:Nautobot:Devices::%s]", err)
				return dev, nil
			}
			d = append(d, item...)
		}
	}
	if len(d) == 0 {
		fmt.Println(text.FgHiYellow.Sprintf("\nUnable to find Nautobot device for %s. Using IP only.", ip.String()))
	}

	dev = n.nbDeviceToSotDevice(d[0])
	dev.IP = ip.String()

	// 4. Finally, return the Nautobot device.
	return dev, nil
}

func deviceSelect(devices []nautobotv1.Device) (nautobotv1.Device, error) {
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
		return nautobotv1.Device{}, err
	}
	for k := range devices {
		if result == devices[k].Name {
			return devices[k], nil
		}
	}
	return nautobotv1.Device{}, fmt.Errorf("unable to determine device from `%s`", result)
}

func (n Nautobot) consolePort(id string) (nautobotv1.ConsolePort, error) {
	p, err := n.ConsoleConnections(&url.Values{"device_id": []string{id}})
	if err != nil {
		return nautobotv1.ConsolePort{}, fmt.Errorf("cfg.Nautobot.ConsoleConnections:%w", err)
	}
	if len(p) == 0 {
		return nautobotv1.ConsolePort{}, fmt.Errorf("no console ports found for device `%s`", id)
	}
	return p[0], nil
}

// deviceByName : Nautobot: Returns a device by name.
func (n Nautobot) getDevice(params *url.Values) (Device, error) {
	// 1. Ignore devices in Offline status
	params.Set("status__n", nautobotv1.StatusOffline)

	// 2. Query Nautobot with the newly built query parameters.
	d, err := n.GetDevices(params)
	if err != nil {
		return Device{}, fmt.Errorf("[Devices::%w]", err)
	}

	// 3. If we found more or less than a single device, we need to prompt
	//    for a single device.
	if len(d) > 1 {
		d[0], err = deviceSelect(d)
		if err != nil {
			return Device{}, fmt.Errorf("[Devices::%w]", err)
		}
	}
	if len(d) == 0 {
		return Device{}, fmt.Errorf("unable to find Nautobot device for `%v`", params)
	}

	if d[0].PrimaryIP == nil {
		return Device{}, fmt.Errorf("`%v` does not have a primary IP assigned in Nautobot", params)
	}

	// 5. Finally, return the device.
	return n.nbDeviceToSotDevice(d[0]), nil
}

func (n Nautobot) nbDeviceToSotDevice(d nautobotv1.Device) Device {
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
		Location: d.Site.Name,
		Status:   d.Status.Value,
		Serial:   d.Serial,
		AssetTag: d.AssetTag,
		Tenant:   d.Tenant.Name,
	}
	if p, err := n.consolePort(d.ID); err == nil {
		ret.Console = Console{
			Hostname: p.ConnectedEndpoint.Device.Name,
			Port:     p.ConnectedEndpoint.Name,
			ID:       p.Device.ID,
		}
	}
	return ret
}
