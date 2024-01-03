// Package version is the version management logic for nbot to make sure we are
// able to manage the version releases.
package version

import (
	"fmt"
	"runtime"
	"time"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/core"
	"github.com/josh-silvas/nbot/core/keyring"
	"github.com/josh-silvas/nbot/nlog"
)

const pluginName = "version"

// Plugin type is used as the command and calling function for each plugin
type Plugin struct {
	core.PluginBase
}

func (p Plugin) Register(c *core.Parser) core.PluginIfc {
	p.Log = nlog.NewWithGroup(pluginName)
	p.C = c.NewCommand(pluginName, "display current version")
	return p
}

func (p Plugin) CMD() *argparse.Command {
	return p.C
}

// pluginFunc function is executed from the nbot caller
func (p Plugin) Func(cfg keyring.Settings) {
	var storedVer ConfigVersion
	key, err := FromConfigFile(cfg)
	if err == nil {
		if storedVer, err = ParseConfigVersion(key.String()); err != nil {
			p.Log.Errorf("ParseConfigVersion(%s)", err)
		}
	}
	fmt.Println(text.FgGreen.Sprintf("NBot: v%s", storedVer.Version.String()))
	fmt.Println(text.FgGreen.Sprintf(" ° Runtime: %s_%s", runtime.GOOS, runtime.GOARCH))
	fmt.Println(text.FgGreen.Sprintf(" ° Version Checked At: %s", storedVer.Timestamp.String()))
	fmt.Println(text.FgGreen.Sprintf(" ° Next Version Check At: %s\n", storedVer.Timestamp.Add(checkInterval*time.Hour)))
}
