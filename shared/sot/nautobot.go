package sot

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/shared/sot/nautobot"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

// Nautobot implementation of the SoT interface.
type Nautobot struct {
	nautobot.Client
}

// NewNautobot : Creates a new Nautobot client.
func NewNautobot(token, nbURL string, opts ...nautobot.Option) (*Nautobot, error) {
	c, err := nautobot.New(token, nbURL, opts...)
	if err != nil {
		return nil, err
	}
	return &Nautobot{Client: *c}, nil
}

// Devices : Nautobot: Returns a list of devices from Nautobot.
func (n *Nautobot) Devices(values *url.Values) ([]Device, error) {
	d, err := n.GetDevices(values)
	if err != nil {
		return nil, err
	}
	devices := make([]Device, 0)
	for i := range d {
		devices = append(devices, Device{
			Hostname: d[i].Name,
			IP:       d[i].PrimaryIP.Address,
			Platform: d[i].Platform.Name,
			Location: d[i].Site.Name,
		})
	}
	return devices, nil
}

// DeviceByName : Nautobot: Returns a device by name.
func (n *Nautobot) DeviceByName(name string) (Device, error) {
	// 1. First, attempt to convert the argD variable to an integer,
	//    if successful, we can assume this is a Nautobot ID.
	params := make(url.Values)
	if _, err := strconv.Atoi(name); err == nil {
		params["id"] = []string{name}
	}

	// 2. Second, if there is still not a Nautobot ID, but the argD variable is
	//    set with something, then we can include this into the name search in Nautobot.
	if _, ok := params["id"]; !ok && name != "" {
		params["name__ic"] = []string{name}
	}

	// 3. Ignore devices in Offline status
	params.Set("status__n", nautobot.StatusOffline)

	// 4. Query Nautobot with the newly built query parameters.
	d, err := n.GetDevices(&params)
	if err != nil {
		return Device{}, fmt.Errorf("[Devices::%w]", err)
	}

	// 4. If we found more or less than a single device, we need to prompt
	//    for a single device.
	if len(d) > 1 {
		d[0], err = deviceSelect(d)
		if err != nil {
			return Device{}, fmt.Errorf("[Devices::%w]", err)
		}
	}
	if len(d) == 0 {
		return Device{}, fmt.Errorf("unable to find Nautobot device for `%s`", name)
	}

	if d[0].PrimaryIP == nil {
		return Device{}, fmt.Errorf("`%s` does not have a primary IP assigned in Nautobot", name)
	}

	// 5. Finally, return the device.
	return Device{
		IP: func() string {
			a := d[0].PrimaryIP.Address
			if strings.Contains(a, "/") {
				a = strings.Split(a, "/")[0]
			}
			return a
		}(),
		Hostname: d[0].Name,
		Platform: d[0].Platform.Name,
		Location: d[0].Site.Name,
	}, nil
}

// DeviceByIP : Nautobot: Returns a device by IP address.
func (n *Nautobot) DeviceByIP(ip net.IP) (Device, error) {
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
	d := make([]nautobot.Device, 0)
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
	dev.Platform = d[0].Platform.Name
	dev.Location = d[0].Site.Name

	// 4. Finally, return the Nautobot device.
	return dev, nil
}

func deviceSelect(devices []nautobot.Device) (nautobot.Device, error) {
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
		return nautobot.Device{}, err
	}
	for k := range devices {
		if result == devices[k].Name {
			return devices[k], nil
		}
	}
	return nautobot.Device{}, fmt.Errorf("unable to determine device from `%s`", result)
}
