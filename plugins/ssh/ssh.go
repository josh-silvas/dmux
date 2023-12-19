package sshinteractive

import (
	"fmt"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/josh-silvas/nbot/shared/connection"
	"github.com/sirupsen/logrus"
)

var (
	argD *string
	argP *string
	argC *string
	argU *string
)

// Plugin function will return an argparse.Command type back to the parent parser
// nolint:typecheck
func Plugin(p *core.Parser) core.Plugin {
	cmd := p.NewCommand("ssh", "Opens an interactive ssh shell.")
	argD = cmd.StringPositional(&argparse.Options{Required: true, Help: "One of DeviceName, IPAddress, DeviceID"})
	argP = cmd.String("p", "port", &argparse.Options{Help: "SSH port to use. Default: 22", Default: "22"})
	argC = cmd.String("c", "command", &argparse.Options{Help: "Optional: Single command to run on the device and exit."})
	argU = cmd.String("u", "user", &argparse.Options{Help: "Optional: Override user credentials to connect to device."})
	return core.Plugin{CMD: cmd, Func: pluginFunc}
}

// pluginFunc function is executed from the caller
func pluginFunc(cfg keyring.Settings) {
	// 1. Check that a device has been passed in
	if strings.TrimSpace(*argD) == "" {
		logrus.Fatalf("A device must be specified! `nbot ssh [DeviceName, IPAddress, DeviceID]`")
	}

	rKey, err := cfg.DeviceAuth()
	if err != nil {
		logrus.Fatalf("RADIUS(%s)", err)
	}

	// 3. If there is a user passed in, attempt to fetch credentials from the
	// keychain, or prompt for password.
	if *argU != "" {
		rKey, err = cfg.UserPassCustom(*argU)
		if err != nil {
			logrus.Fatalf("DeviceAuth(%s)", err)
		}
	}

	// 5. Initialize a device from NetBox or from DNS.
	d, err := connection.NewDevice(*argD)
	if err != nil {
		logrus.Fatalf("Device(%s)", err)
	}

	if *argC != "" {
		out, err := d.RunCommand(*argC, rKey.Username, rKey.Password(), *argP)
		if err != nil {
			logrus.Fatalf("RunCommands(%s)", err)
		}
		fmt.Println(string(out))
		return
	}

	fmt.Println(text.FgHiCyan.Sprintf("\nLogging into %s: %s [%s] ", d.ID, d.Hostname, d.Status.Label))
	if d.Comments != "" {
		fmt.Println(text.FgHiYellow.Sprintf("Device has comments:\n%s", d.Comments))
	}

	// 6. Launch interactive shell.
	if err := d.InteractiveShell(rKey.Username, rKey.Password(), *argP); err != nil {
		logrus.Fatalf("InteractiveShell(%s)", err)
	}
}
