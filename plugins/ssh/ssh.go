package sshinteractive

import (
	"fmt"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/josh-silvas/nbot/nlog"
	"github.com/josh-silvas/nbot/shared/connection"
	"github.com/josh-silvas/nbot/shared/sot"
)

const pluginName = "ssh"

// Plugin type is used as the command and calling function for each plugin
type Plugin struct {
	core.PluginBase
	argD *string
	argP *string
	argC *string
	argU *string
}

func (p Plugin) Register(c *core.Parser) core.PluginIfc {
	p.Log = nlog.NewWithGroup(pluginName)
	p.C = c.NewCommand("ssh", "Opens an interactive ssh shell.")
	p.argD = p.C.StringPositional(&argparse.Options{Required: true, Help: "One of DeviceName, IPAddress, DeviceID"})
	p.argP = p.C.String("p", "port", &argparse.Options{Help: "SSH port to use. Default: 22", Default: "22"})
	p.argC = p.C.String("c", "command", &argparse.Options{Help: "Optional: Single command to run on the device and exit."})
	p.argU = p.C.String("u", "user", &argparse.Options{Help: "Optional: Override user credentials to connect to device."})
	return p
}

func (p Plugin) CMD() *argparse.Command {
	return p.C
}

func (p Plugin) Func(cfg keyring.Settings) {
	// 1. Check that a device has been passed in
	if strings.TrimSpace(*p.argD) == "" {
		p.Log.Fatal("A device must be specified! `nbot ssh [DeviceName, IPAddress, DeviceID]`")
	}

	rKey, err := cfg.DeviceAuth()
	if err != nil {
		p.Log.Fatalf("RADIUS(%s)", err)
	}

	// 3. If there is a user passed in, attempt to fetch credentials from the
	// keychain, or prompt for password.
	if *p.argU != "" {
		rKey, err = cfg.UserPassCustom(*p.argU)
		if err != nil {
			p.Log.Fatalf("DeviceAuth(%s)", err)
		}
	}

	nb, err := sot.New(cfg)
	if err != nil {
		p.Log.Fatalf("New(%s)", err)
	}

	// 5. Initialize a device from NetBox or from DNS.
	d, err := connection.NewDevice(nb, *p.argD)
	if err != nil {
		p.Log.Fatalf("Device(%s)", err)
	}

	if *p.argC != "" {
		out, err := d.RunCommand(*p.argC, rKey.Username, rKey.Password(), *p.argP)
		if err != nil {
			p.Log.Fatalf("RunCommands(%s)", err)
		}
		fmt.Println(string(out))
		return
	}

	fmt.Println(text.FgHiCyan.Sprintf("\nLogging into %s: %s [%s] ", d.ID, d.Hostname, d.Status))
	if d.Comments != "" {
		fmt.Println(text.FgHiYellow.Sprintf("Device has comments:\n%s", d.Comments))
	}

	// 6. Launch interactive shell.
	if err := d.InteractiveShell(rKey.Username, rKey.Password(), *p.argP); err != nil {
		p.Log.Fatalf("InteractiveShell(%s)", err)
	}
}
