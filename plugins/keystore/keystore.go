package keystore

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/josh-silvas/nbot/nlog"
)

const pluginName = "keystore"

// Plugin type is used as the command and calling function for each plugin
type Plugin struct {
	core.PluginBase
	argR *string
	argV *string
}

// Register : register the plugin with the parser
func (p Plugin) Register(c *core.Parser) core.PluginIfc {
	p.Log = nlog.NewWithGroup(pluginName)
	p.C = c.NewCommand(pluginName, "Return current keychain data.")
	p.argR = p.C.String("", "reset", &argparse.Options{Help: "Reset keyring creds by service-name"})
	p.argV = p.C.String("", "view", &argparse.Options{Help: "View keyring creds by service-name"})
	return p
}

// CMD : return the command for the plugin
func (p Plugin) CMD() *argparse.Command {
	return p.C
}

// Func : pluginFunc function is executed from the nbot caller
func (p Plugin) Func(cfg keyring.Settings) {
	if *p.argR != "" {
		if err := cfg.Delete(*p.argR); err != nil {
			p.Log.Fatalf("cfg.Delete(%s)", err)
		}
		return
	}

	if *p.argV != "" {
		r, err := cfg.Get(*p.argV)
		if err != nil {
			p.Log.Fatalf("cfg.Get(%s)", err)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{*p.argV, "Key"})
		t.AppendRows([]table.Row{{r.Username, r.Password()}})
		t.SetStyle(table.StyleColoredBlackOnGreenWhite)
		t.Render()
	}
}
