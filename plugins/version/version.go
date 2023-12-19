// Package version is the version management logic for nbot to make sure we are
// able to manage the version releases.
package version

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/sirupsen/logrus"
)

// Plugin function will return an argparse.Command type back to the parent parser
// nolint:typecheck
func Plugin(p *core.Parser) core.Plugin {
	cmd := p.NewCommand("version", "display current version")
	return core.Plugin{CMD: cmd, Func: pluginFunc}
}

// pluginFunc function is executed from the nbot caller
func pluginFunc(cfg keyring.Settings) {
	var storedVer ConfigVersion
	key, err := FromConfigFile(cfg)
	if err == nil {
		if storedVer, err = ParseConfigVersion(key.String()); err != nil {
			logrus.Error(err)
		}
	}
	fmt.Println(text.FgGreen.Sprintf("NBot: v%s", storedVer.Version.String()))
	fmt.Println(text.FgGreen.Sprintf(" ° Runtime: %s_%s", runtime.GOOS, runtime.GOARCH))
	fmt.Println(text.FgGreen.Sprintf(" ° Version Checked At: %s", storedVer.Timestamp.String()))
	fmt.Println(text.FgGreen.Sprintf(" ° Next Version Check At: %s\n", storedVer.Timestamp.Add(checkInterval*time.Hour)))
}
