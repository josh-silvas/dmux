package connection

import (
	"net"
	"strconv"
	"strings"

	"github.com/josh-silvas/dmux/pkg/sot"
)

// Device : Nested device struct, so we can attach methods.
type Device struct {
	sot.Device
}

// NewDevice : Helper method to fetch a device from NautobotV1 using a
// set of different methods, explained below.
func NewDevice(c sot.SoT, arg string) (Device, error) {
	// 1. Check if the value passed in is a valid IPv4 address. If
	//    so, then get the device by IP address.
	if ip := net.ParseIP(arg); ip != nil {
		device, err := c.GetDevice(sot.ByIP, ip)
		if err != nil {
			return Device{}, err
		}
		return Device{device}, nil
	}

	// 2. Next, check if the value passed in can be resolved via DNS. If
	//    so, take the DNS resolved IP address and fetch by IP.
	if ip := dnsLookup(arg); ip != nil {
		device, err := c.GetDevice(sot.ByIP, ip)
		if err != nil {
			return Device{}, err
		}
		return Device{device}, nil
	}

	// 3. Finally, attempt to fetch the device by a partial case-insensitive
	//    device name match in NautobotV1.
	device, err := c.GetDevice(sot.ByName, arg)
	if err != nil {
		return Device{}, err
	}
	return Device{device}, nil
}

func dnsLookup(s string) net.IP {
	// 1. Perform a DNS lookup of the hostname string
	//    to determine if it's a valid DNS entry.
	ips, err := net.LookupIP(s)
	if err != nil {
		return nil
	}

	// 2. If we get a valid DNS entry, search for the first
	//    available IPv4 DNS entry and return that one.
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip
		}
	}
	return nil
}

// HostPort : Helper method on the Device struct used to gather the
// IP Address/SSH port combination for this device.
func (d Device) HostPort(port string) string {
	a := d.IP // nolint: typecheck
	if strings.Contains(a, "/") {
		a = strings.Split(a, "/")[0]
	}
	if _, err := strconv.Atoi(port); err != nil {
		port = "22"
	}
	return net.JoinHostPort(a, port)
}

// InteractiveShell func will kick off an interactive ssh shell
func (d Device) InteractiveShell(user, pass, port string) error {
	c, err := NewDialer(d.HostPort(port), user, pass)
	if err != nil {
		l.Fatalf("New::%s", err)
	}

	// Create Session
	session, err := c.Client.NewSession()
	if err != nil {
		l.Fatalf("Client.NewSession::%s", err)
	}

	// Set a terminal log for this session.
	c.SetLog(d.IP) // nolint: typecheck

	// Start ssh shell
	return c.Shell(session)
}

// RunCommand func will kick off an ssh session and run arbitrary commands on a device.
func (d Device) RunCommand(cmd string, user, pass, port string) ([]byte, error) {
	c, err := NewDialer(d.HostPort(port), user, pass)
	if err != nil {
		return nil, err
	}
	// Create Session
	session, err := c.Client.NewSession()
	if err != nil {
		return nil, err
	}

	defer func() {
		if e := session.Close(); e != nil {
			err = e
		}
	}()
	return session.Output(cmd)
}
