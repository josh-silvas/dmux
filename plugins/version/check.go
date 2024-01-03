package version

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/go-ini/ini"
	"github.com/google/go-github/github"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/internal/core"
	"github.com/josh-silvas/nbot/internal/keyring"

	"github.com/Masterminds/semver"
)

const (
	checkInterval = 48
	owner         = "josh-silvas"
)

// Check function is executed from the nbot caller
func Check(cfg keyring.Settings) error {
	runningVer, err := SemVer(cfg.Meta["buildVersion"])
	if err != nil {
		return fmt.Errorf("version check failed: %s", err)
	}

	key, err := FromConfigFile(cfg)
	if err != nil {
		return fmt.Errorf("version check failed: %s", err)
	}
	storedVer, err := ParseConfigVersion(key.String())
	if err != nil {
		return fmt.Errorf("version check failed: %s", err)
	}
	key.SetValue(ConfigVersion{Version: runningVer, Timestamp: time.Now()}.String())
	if err = cfg.File.SaveTo(cfg.Source); err != nil {
		return fmt.Errorf("version check failed: %s", err)
	}

	// Here we are checking if the timestamp on the cached version is more than
	// 12 hours old. If it's not then we can just exit here.
	if storedVer.Timestamp.After(time.Now().Add(-checkInterval * time.Hour)) {
		return nil
	}

	apiVer, err := FromGitHub()
	if err != nil {
		return fmt.Errorf("version check failed: %s", err)
	}
	if runningVer.LessThan(apiVer) {
		versionPrint(runningVer, apiVer)
	}
	return nil
}

// FromGitHub : Fetches the latest version from GitHub.
func FromGitHub() (*semver.Version, error) {
	latestVer, err := semver.NewVersion("0.0.0")
	if err != nil {
		return latestVer, err
	}
	client := github.NewClient(nil)

	opt := &github.ListOptions{Page: 1, PerPage: 10}
	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, core.AppName, opt)
	if err != nil {
		return latestVer, err
	}
	for _, r := range releases {
		thisVer, err := semver.NewVersion(strings.ReplaceAll(r.GetTagName(), "v", ""))
		if err != nil {
			continue
		}
		if thisVer.Prerelease() != "" {
			continue
		}
		if thisVer.GreaterThan(latestVer) {
			latestVer = thisVer
		}
	}

	if err != nil {
		return latestVer, err
	}

	return latestVer, nil
}

// FromConfigFile : From gokeys.Settings type to open the config file and retrieve the cached version.
func FromConfigFile(cfg keyring.Settings) (*ini.Key, error) {
	sec, err := cfg.File.GetSection("nbot")
	// If there is an error retrieving the section, it likely is not created yet.
	// attempt to create the section
	if err != nil {
		sec, err = cfg.File.NewSection("nbot")
		if err != nil {
			return nil, err
		}

	}
	key, err := sec.GetKey("version")
	if err != nil {
		key, err = sec.NewKey("version", "v0.0.0::0001-01-01T00:00:00Z")
		if err != nil {
			return nil, err
		}
		key.Comment = "local cache of the nbot version, really for the timestamp"
	}
	return key, nil
}

// ConfigVersion : the parsed type from the configuration file stored locally
type ConfigVersion struct {
	Version   *semver.Version
	Timestamp time.Time
}

// String method set onto the ConfigVersion type will convert the type into a
// value that is expected and consistent so that it can be parsed later.
func (c ConfigVersion) String() string {
	// As a standard, we are going to use time.RFC3339 as the timestamp storage format.
	return fmt.Sprintf("%s::%s", c.Version, c.Timestamp.Format(time.RFC3339))
}

// ParseConfigVersion will take a string value and attempt to parse is into a ConfigVersion type.
// If the string is set to its null value, then return an empty type and no error as
// this would be the accurate representation of the parsed item.
func ParseConfigVersion(c string) (ConfigVersion, error) {
	// Do not error if the string is empty, simply return
	// the empty value of the ConfigVersion type.
	if c == "" {
		return ConfigVersion{}, nil
	}
	arr := strings.Split(c, "::")
	if len(arr) != 2 {
		return ConfigVersion{}, fmt.Errorf("unable to parse version from cfg: %s", c)
	}
	ts, err := time.Parse(time.RFC3339, strings.TrimSpace(arr[1]))
	if err != nil {
		return ConfigVersion{}, err
	}
	ver, err := semver.NewVersion(arr[0])
	if err != nil {
		return ConfigVersion{}, err
	}
	return ConfigVersion{Version: ver, Timestamp: ts}, nil
}

// SemVer is a helper to convert goreleaser/git tags with semver tags.
func SemVer(s string) (*semver.Version, error) {
	if s == "" {
		s = "0.0.0"
	}
	ver, err := semver.NewVersion(strings.TrimPrefix(s, "v"))
	if err != nil {
		return nil, fmt.Errorf("version check failed: %w:%s", err, s)
	}
	return ver, nil
}

// versionPrint is used to print info to terminal if the user needs to be notified of a new or different running version
func versionPrint(running, current *semver.Version) {
	fmt.Println(text.FgHiCyan.Sprintf("Upgrade available (%s running, %s available). Install with:", running, current))
	fmt.Println(text.FgYellow.Sprintf("   >> nbot upgrade"))
	fmt.Println(text.FgYellow.Sprintf("   >> - OR - "))
	switch runtime.GOOS {
	case "linux":
		fmt.Println(text.FgHiYellow.Sprintf("   >> curl -O https://github.com/josh-silvas/nbot/releases/%s/nbot_64-bit.deb "+
			"&& sudo dpkg -i nbot_64-bit.deb", current.String()))
	case "darwin":
		fmt.Println(text.FgYellow.Sprintf("   >> brew update && brew upgrade nbot"))
	default:
		fmt.Println(text.FgYellow.Sprintf("Unknown OS, check https://github.com/josh-silvas/nbot for install options"))
	}
	fmt.Println(text.FgYellow.Sprintf("You will be notified in %d hours if you have not upgraded.", checkInterval))
}
