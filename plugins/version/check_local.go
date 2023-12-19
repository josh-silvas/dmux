package version

import (
	"github.com/go-ini/ini"
	"github.com/josh-silvas/nbot/core/keyring"
)

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
