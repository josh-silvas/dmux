package sot

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/manifoldco/promptui"
	"github.com/netbox-community/go-netbox/v3/netbox/client"
	"github.com/netbox-community/go-netbox/v3/netbox/client/dcim"
	"github.com/netbox-community/go-netbox/v3/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// Netbox implementation of the SoT interface.
type Netbox struct {
	client.NetBoxAPI
}

// NewNetbox : Creates a new Netbox client.
func NewNetbox(token, nbURL string) (*Netbox, error) {
	t := transport.New(nbURL, client.DefaultBasePath, []string{"https"})
	t.DefaultAuthentication = transport.APIKeyAuth(
		"Authorization",
		"header",
		fmt.Sprintf("Token %v", token),
	)
	c := client.New(t, strfmt.Default)
	return &Netbox{*c}, nil
}

// GetDevice : NautobotV1: Returns a single device from NautobotV1.
func (n Netbox) GetDevice(method string, value any) (Device, error) {
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

// deviceByNameOrID : NautobotV1: Returns a device by name or UUIDv4 ID.
func (n Netbox) deviceByNameOrID(name string) (Device, error) {
	// 1. First, attempt to convert the argD variable to an integer,
	//    if successful, we can assume this is a Netbox ID.
	params := new(dcim.DcimDevicesListParams)

	if _, err := strconv.Atoi(name); err == nil {
		params.ID = &name
	}

	// 2. Second, if there is still not a Netbox ID, but the name variable is
	//    set with something, then we can include this into the name search in NautobotV1.
	if params.ID == nil && name != "" {
		params.NameIc = &name
	}

	return n.getDevice(params)
}

// deviceBySerial : Netbox: Returns a device by serial number.
func (n Netbox) deviceBySerial(serial string) (Device, error) {
	params := new(dcim.DcimDevicesListParams)
	params.SerialIe = &serial
	return n.getDevice(params)
}

// deviceByMac : Netbox: Returns a device by mac address.
func (n Netbox) deviceByMac(mac string) (Device, error) {
	params := new(dcim.DcimDevicesListParams)
	params.MacAddressIc = &mac
	return n.getDevice(params)
}

// deviceByIP : Netbox: Returns a device by IP address.
func (n Netbox) deviceByIP(ip net.IP) (Device, error) {
	// 1. Create the device object, we don't want to fail out
	//    in there's a NautobotV1 error, since we already have what we
	//    need to connect.
	dev := Device{
		IP:       ip.String(),
		Hostname: ip.String(),
	}
	// 2. Fetch all the devices from NautobotV1 that match this IP address.
	ips, err := n.Ipam.IpamIPAddressesList(&ipam.IpamIPAddressesListParams{Address: &dev.IP}, nil)
	if err != nil {
		l.Errorf("[COMERR:Netbox:IPAddresses::%s]", err)
		return dev, nil
	}

	type AssignedObjectInterface struct {
		Display string `json:"display"`
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Device  struct {
			Display string `json:"display"`
			ID      int    `json:"id"`
			URL     string `json:"url"`
			Name    string `json:"name"`
		} `json:"device,omitempty"`
		VirtualMachine struct {
			Display string `json:"display"`
			ID      int    `json:"id"`
			URL     string `json:"url"`
			Name    string `json:"name"`
		} `json:"virtual_machine,omitempty"`
		Name string `json:"name"`
	}

	// 3. For each IP address, query the first match found for a device
	//    within NautobotV1.
	d := make([]models.DeviceWithConfigContext, 0)
	for _, i := range ips.GetPayload().Results {
		if i.AssignedObject == nil {
			continue
		}
		if i.AssignedObject != "" {
			var obj AssignedObjectInterface
			byteData, err := json.Marshal(i.AssignedObject)
			if err != nil {
				l.Errorf("[COMERR:Netbox:IPAddresses::%s]", err)
				return dev, nil
			}
			if err := json.Unmarshal(byteData, &obj); err != nil {
				l.Errorf("[COMERR:Netbox:IPAddresses::%s]", err)
				return dev, nil
			}
			id := func() int64 {
				if obj.Device.ID != 0 {
					return int64(obj.Device.ID)
				}
				return int64(obj.VirtualMachine.ID)
			}()
			item, err := n.Dcim.DcimDevicesRead(&dcim.DcimDevicesReadParams{ID: id}, nil)
			if err != nil {
				l.Errorf("[COMMERR:NautobotV1:Devices::%s]", err)
				return dev, nil
			}
			d = append(d, *item.Payload)
		}
	}
	if len(d) == 0 {
		fmt.Println(text.FgHiYellow.Sprintf("\nUnable to find NautobotV1 device for %s. Using IP only.", ip.String()))
	}

	dev = n.nbDeviceToSotDevice(d[0])
	dev.IP = ip.String()

	// 4. Finally, return the NautobotV1 device.
	return dev, nil
}

func netBoxDeviceSelect(devices []*models.DeviceWithConfigContext) (models.DeviceWithConfigContext, error) {
	prompt := promptui.Select{
		Label: "Multiple devices found. Select",
		Items: func() []string {
			r := make([]string, 0)
			for k := range devices {
				r = append(r, *devices[k].Name)
			}
			return r
		}(),
	}

	_, result, err := prompt.Run()
	if err != nil {
		return models.DeviceWithConfigContext{}, err
	}
	for k := range devices {
		if result == *devices[k].Name {
			return *devices[k], nil
		}
	}
	return models.DeviceWithConfigContext{}, fmt.Errorf("unable to determine device from `%s`", result)
}

func (n Netbox) consolePort(id int) (models.ConsolePort, error) { // nolint: unused
	strID := strconv.Itoa(id)
	p, err := n.Dcim.DcimConsolePortsList(&dcim.DcimConsolePortsListParams{DeviceID: &strID}, nil)
	if err != nil {
		return models.ConsolePort{}, fmt.Errorf("cfg.Netbox.DcimConsolePortsList:%w", err)
	}
	if len(p.Payload.Results) == 0 {
		return models.ConsolePort{}, fmt.Errorf("no console ports found for device `%d`", id)
	}
	return *p.Payload.Results[0], nil
}

// deviceByName : Netbox: Returns a device by name.
func (n Netbox) getDevice(params *dcim.DcimDevicesListParams) (Device, error) {
	// 1. Ignore devices in Offline status
	offline := "offline"
	params.Statusn = &offline
	params.Context = context.Background()

	// 2. Query NautobotV1 with the newly built query parameters.
	d, err := n.Dcim.DcimDevicesList(params, nil)
	if err != nil {
		return Device{}, fmt.Errorf("[Devices::%w]", err)
	}

	if *d.Payload.Count == 0 {
		return Device{}, fmt.Errorf("unable to find Netbox device for `%v`", params)
	}

	device := *d.Payload.Results[0]
	// 3. If we found more or less than a single device, we need to prompt
	//    for a single device.
	if len(d.Payload.Results) > 1 {
		device, err = netBoxDeviceSelect(d.Payload.Results)
		if err != nil {
			return Device{}, fmt.Errorf("[Devices::%w]", err)
		}
	}

	if device.PrimaryIP == nil {
		return Device{}, fmt.Errorf("`%v` does not have a primary IP assigned in NautobotV1", params)
	}

	// 5. Finally, return the device.
	return n.nbDeviceToSotDevice(device), nil
}

func (n Netbox) nbDeviceToSotDevice(d models.DeviceWithConfigContext) Device {
	ret := Device{
		Hostname: *d.Name,
		IP: func() string {
			if d.PrimaryIP == nil {
				return ""
			}
			a := *d.PrimaryIP.Address
			if strings.Contains(a, "/") {
				a = strings.Split(a, "/")[0]
			}
			return a
		}(),
		Platform: func() string {
			if d.DeviceType == nil {
				return ""
			}
			return *d.DeviceType.Model
		}(),
		Location: func() string {
			if d.Site == nil {
				return ""
			}
			return *d.Site.Name
		}(),
		Status: func() string {
			if d.Status == nil {
				return ""
			}
			return *d.Status.Label
		}(),
		Serial: d.Serial,
		AssetTag: func() string {
			if d.AssetTag == nil {
				return ""
			}
			return *d.AssetTag
		}(),
		Tenant: func() string {
			if d.Tenant == nil {
				return ""
			}
			return *d.Tenant.Name
		}(),
	}
	return ret
}
