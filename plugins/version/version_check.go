package version

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/core/keyring"

	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"
)

const checkInterval = 24

// Check function is executed from the nbot caller
func Check(cfg keyring.Settings) error {
	runningVer := SemVer(cfg.Meta["buildVersion"])
	key, err := FromConfigFile(cfg)
	if err != nil {
		return fmt.Errorf("version check failed: %w", err)
	}
	storedVer, err := ParseConfigVersion(key.String())
	if err != nil {
		return fmt.Errorf("version check failed: %w", err)
	}
	key.SetValue(ConfigVersion{Version: runningVer, Timestamp: time.Now()}.String())
	if err = cfg.File.SaveTo(cfg.Source); err != nil {
		return fmt.Errorf("version check failed: %w", err)
	}

	// Here we are checking if the timestamp on the cached version is more than
	// 12 hours old. If it's not then we can just exit here.
	if storedVer.Timestamp.After(time.Now().Add(-checkInterval * time.Hour)) {
		return nil
	}

	apiVer, err := FromArtifactory("")
	if err != nil {
		return fmt.Errorf("version check failed: %w", err)
	}
	if runningVer.LessThan(apiVer) {
		versionPrint(runningVer, apiVer)
	}
	return nil
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
func SemVer(s string) *semver.Version {
	if s == "" {
		s = "0.0.0"
	}
	ver, err := semver.NewVersion(strings.TrimPrefix(s, "v"))
	if err != nil {
		logrus.Errorf("Version check failed: %s:%s", err, s)
	}
	return ver
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
