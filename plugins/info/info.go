package info

import (
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/josh-silvas/nbot/shared"
	"github.com/josh-silvas/nbot/shared/sot/nautobot"
	"github.com/sirupsen/logrus"
)

var (
	argD  *string
	argT  *string
	argS  *string
	argSe *string
	argM  *string
	argV  *bool
)

// Plugin function will return an argparse.Command type back to the parent parser
// nolint:typecheck
func Plugin(p *core.Parser) core.Plugin {
	cmd := p.NewCommand("info", "Gathers device information.")
	argD = cmd.StringPositional(&argparse.Options{Required: true, Help: "One of DeviceName, IPAddress, DeviceID"})
	argT = shared.ArgString(cmd, "tag", "Devices with a tag")
	argS = shared.ArgString(cmd, "site", "Devices within a given site name.")
	argM = shared.ArgString(cmd, "mac-address", "Devices within a given mac-address.")
	argSe = cmd.String("", "serial", &argparse.Options{Help: "Devices by serial number."})
	argV = cmd.Flag("v", "verbose", &argparse.Options{Help: "Prints additional information [tags, manufacturer, etc]."})
	return core.Plugin{CMD: cmd, Func: pluginFunc}
}

// pluginFunc function is executed from the caller
func pluginFunc(cfg keyring.Settings) {
	// 1. Check that a device has been passed in
	if strings.TrimSpace(*argD) == "" {
		logrus.Fatalf("A device must be specified! `nbot info [DeviceName, IPAddress, DeviceID]`")
	}

	key, err := cfg.Nautobot()
	if err != nil {
		logrus.Fatalf("cfg.Nautobot:VaultRead:%s", err)
	}
	nbOpts := make([]nautobot.Option, 0)
	nb, err := nautobot.New(key.Password(), nbOpts...)
	if err != nil {
		logrus.Fatalf("nautobot.New:%s", err)
	}

	// Separate handle if an IP address is passed in as a query param
	if ip := net.ParseIP(*argD); ip != nil {
		ips, err := nb.GetIPAddresses(&url.Values{"address": []string{ip.String()}})
		if err != nil {
			logrus.Fatalf("cfg.Nautobot:IPAddresses:%s", err)
		}
		devices := make([]shared.DeviceAggregate, 0)
		for i := range ips {
			if ips[i].AssignedObject.Device.ID != "" {
				nbd, err := shared.GetDeviceAggregates(nb, url.Values{"id": []string{ips[i].AssignedObject.Device.ID}})
				if err != nil {
					logrus.Fatalf("cfg.Nautobot:Devices:%s", err)
				}
				devices = append(devices, nbd...)
			}
		}
		shared.PrintDevices(devices, *argV)
		return
	}

	// Logic to handle device related query params.
	params := make(url.Values)
	if *argT != "" {
		params["tag"] = []string{*argT}
	}
	if *argS != "" {
		params["site"] = []string{*argS}
	}
	if *argSe != "" {
		params["serial"] = []string{*argSe}
	}
	if *argM != "" {
		if hw, err := net.ParseMAC(*argM); err == nil {
			params["mac_address__ie"] = []string{hw.String()}
		} else {
			logrus.Errorf("unable to parse mac-address [%s]", *argM)
		}
	}
	if _, err := strconv.Atoi(*argD); err == nil {
		params["id"] = []string{*argD}
	}
	if _, ok := params["id"]; !ok && *argD != "" {
		params["name__ic"] = []string{*argD}
	}
	if len(params) == 0 {
		logrus.Fatalln("a query must be specified to retrieve devices from Nautobot")
	}
	devices, err := shared.GetDeviceAggregates(nb, params)
	if err != nil {
		logrus.Fatalf("cfg.Nautobot.Devices:%s", err)
	}
	shared.PrintDevices(devices, *argV)
}
