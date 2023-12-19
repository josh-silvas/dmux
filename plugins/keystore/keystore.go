package keystore

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/shared/keyring"
	"github.com/sirupsen/logrus"
)

var (
	argR *string
	argV *string
)

// Plugin : Instantiates the core.Plugin addition of this app.
// nolint:typecheck
func Plugin(p *core.Parser) core.Plugin {
	cmd := p.NewCommand("keystore", "Return current keychain data.")
	argR = cmd.String("", "reset", &argparse.Options{Help: "Reset keyring creds by service-name"})
	argV = cmd.String("", "view", &argparse.Options{Help: "View keyring creds by service-name"})
	return core.Plugin{CMD: cmd, Func: pluginFunc}
}

// pluginFunc is the main entry point function and will gather each token type from the keyring service
func pluginFunc(cfg keyring.Settings) {
	if *argR != "" {
		if err := cfg.Delete(*argR); err != nil {
			logrus.Fatal(err)
		}
		return
	}

	if *argV != "" {
		r, err := cfg.Get(*argV)
		if err != nil {
			logrus.Fatal(err)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{*argV, "Key"})
		t.AppendRows([]table.Row{{r.Username, r.Password()}})
		t.SetStyle(table.StyleColoredBlackOnGreenWhite)
		t.Render()
	}
}
