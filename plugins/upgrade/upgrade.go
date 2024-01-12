package upgrade

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/akamensky/argparse"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/internal/core"
	"github.com/josh-silvas/nbot/internal/keyring"
	"github.com/josh-silvas/nbot/internal/nlog"
	"github.com/josh-silvas/nbot/plugins/version"
)

const (
	msgFailed  = "%s\n\n   >> Unable to upgrade to %s. ☝️️️See above for errors.️️☝️"
	msgSuccess = "%s\n\n   >> NBot successfully upgraded!"
	linuxCmd1  = "curl -O https://github.com/josh-silvas/nbot/releases/%s/nbot_64-bit.deb"
	linuxCmd2  = "sudo dpkg -i nbot_64-bit.deb"
	osxCmd1    = "brew update"
	osxCmd2    = "brew upgrade nbot"
)

const pluginName = "upgrade"

// Plugin type is used as the command and calling function for each plugin
type Plugin struct {
	core.PluginBase
}

// Register : registers the plugin with the parser
func (p Plugin) Register(c *core.Parser) core.Plugin {
	p.Log = nlog.NewWithGroup(pluginName)
	p.C = c.NewCommand(pluginName, "Attempt an upgrade of NBot.")
	return p
}

// CMD : returns the command for the plugin
func (p Plugin) CMD() *argparse.Command {
	return p.C
}

// Func : function that will be executed from the nbot caller
func (p Plugin) Func(cfg keyring.Settings) {
	runningVer, err := version.SemVer(cfg.Meta["buildVersion"])
	if err != nil {
		p.Log.Fatalf("[NBot Upgrade] %s", err)
	}
	key, err := version.FromConfigFile(cfg)
	if err != nil {
		p.Log.Fatalf("[NBot Upgrade] %s", err)
	}
	key.SetValue(version.ConfigVersion{Version: runningVer, Timestamp: time.Now()}.String())
	if err = cfg.File.SaveTo(cfg.Source); err != nil {
		p.Log.Fatalf("[NBot Upgrade] Version check failed. %s", err)
	}

	apiVer, err := version.FromGitHub()
	if err != nil {
		p.Log.Fatalf("[NBot Upgrade] Version check failed. %s", err)
	}

	if !runningVer.LessThan(apiVer) {
		fmt.Println(text.FgHiGreen.Sprintf(
			"   >> ❤️ Thanks for checking in, but you're already on the latest version v%s ❤️", runningVer),
		)
		return
	}
	switch runtime.GOOS {
	case "linux":
		fmt.Println(text.FgHiYellow.Sprintf("   >> Pulling latest .deb NBot package..."))
		if out, err := exec.Command("/bin/bash", "-c", linuxCmd1).CombinedOutput(); err != nil {
			fmt.Println(text.FgHiRed.Sprintf(msgFailed, string(out), apiVer.String()))
			os.Exit(1)
		}
		fmt.Println(text.FgHiYellow.Sprintf("   >> Installing NBot v%s...", apiVer.String()))
		out, err := exec.Command("/bin/bash", "-c", linuxCmd2).CombinedOutput()
		if err != nil {
			fmt.Println(text.FgHiRed.Sprintf(msgFailed, string(out), apiVer.String()))
			os.Exit(1)
		}
		fmt.Println(text.FgHiGreen.Sprintf(msgSuccess, string(out)))
	case "darwin":
		fmt.Println(text.FgHiYellow.Sprintf("   >> Running brew update, be back in a few... ☕️"))
		if out, err := exec.Command("/bin/bash", "-c", osxCmd1).CombinedOutput(); err != nil {
			fmt.Println(text.FgHiRed.Sprintf(msgFailed, string(out), apiVer.String()))
			os.Exit(1)
		}
		fmt.Println(text.FgHiYellow.Sprintf("   >> Upgrading NBot to v%s... 🛠", apiVer.String()))
		out, err := exec.Command("/bin/bash", "-c", osxCmd2).CombinedOutput()
		if err != nil {
			fmt.Println(text.FgHiRed.Sprintf(msgFailed, string(out), apiVer.String()))
			os.Exit(1)
		}
		fmt.Println(text.FgHiGreen.Sprintf(msgSuccess, string(out)))
	default:
		p.Log.Error("Unknown OS, check https://github.com/josh-silvas/nbot for install/upgrade options")
	}
}
