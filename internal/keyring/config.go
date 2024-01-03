package keyring

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/keyring"
	"github.com/go-ini/ini"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

const (
	// ConfigPath : Base directory for NBot specific files.
	ConfigPath = ".config/nbot"

	// SettingsFile : Name of the file in the ConfigPath directory that contains
	// nbot settings.
	SettingsFile = "settings.ini"

	// KeyChainName declares a separate keychain as to separate from other keychains
	// that could possibly sync to iCloud or other devices. This keychain will not be able
	// to synchronize
	KeyChainName = "NBot-Keyring"

	// Split : Special set of characters to act as a delimiter between
	// the user and password.
	Split = "!***!"

	// SvcBase : The base string to identify a service in the keyring.
	SvcBase = "com.keyring.nbot"

	fileBackend = "~/.local/share/keyrings/"

	keychainCMD = "/usr/bin/security"

	defaultExpire = 720 * time.Hour // One month
)

// Settings type is the structure representation of
// the keyring ini profile held in the .config directory
type Settings struct {
	User   string
	Key    map[string]keyring.Keyring
	pin    int
	File   *ini.File
	Source string
	Test   bool
	Meta   map[string]string
}

// CreateIfNotExist function will create the directory and file for the config
// if it does not already exist.  If this info already exist, it will do nothing.
func CreateIfNotExist(homeDir string) error {
	path := fmt.Sprintf("%s/%s", homeDir, ConfigPath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		l.Infof("directory does not exist, creating directory at ", path)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, SettingsFile)); os.IsNotExist(err) {
		l.Infof("file does not exist, creating file ", SettingsFile, path)
		file, err := os.Create(fmt.Sprintf("%s/%s", path, SettingsFile))
		defer func() {
			if e := file.Close(); e != nil {
				closeErr := e.Error()
				err = fmt.Errorf("%w:%s", err, closeErr)
			}
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

// GetConfig function takes a home directory path or none to use the user profile directory, and
// loads the ini file into a Settings structure and returns back the loaded config.
func GetConfig(homeDir string) (Settings, error) {
	if homeDir == "" {
		homeDir = userProfile.HomeDir
	}
	var (
		err      error
		settings Settings
	)
	settings.Source = fmt.Sprintf("%s/%s/%s", homeDir, ConfigPath, SettingsFile)

	settings.File, err = ini.InsensitiveLoad(settings.Source)
	if err != nil {
		l.Errorf("error fetching GetConfig:ini.InsensitiveLoad: %s", err)
		return settings, err
	}
	if err := settings.loadBaseSection(settings.File); err != nil {
		return settings, err
	}

	if err := settings.File.SaveTo(settings.Source); err != nil {
		return settings, err
	}
	return settings, nil
}

// KeyFromSection : Helper method to fetch a value from the settings.ini file for a given section/key
func (s *Settings) KeyFromSection(section, key string, def func() (string, error)) (*ini.Key, error) {
	sec, err := s.File.GetSection(section)
	// If there is an error retrieving the section, it likely is not created yet.
	// attempt to create the section
	if err != nil {
		logrus.Infof("Section [%s] not found in %s/%s, creating...", section, ConfigPath, SettingsFile)
		sec, err = s.File.NewSection(section)
		if err != nil {
			return nil, err
		}
	}
	value, err := sec.GetKey(key)
	if err != nil {
		if def == nil {
			return nil, fmt.Errorf("key not found: %s->%s", section, key)
		}
		newVal, err := def()
		if err != nil {
			return nil, err
		}
		value, err = sec.NewKey(key, newVal)
		if err != nil {
			return nil, err
		}
		value.Comment = fmt.Sprintf("Local value for %s->%s stored in NBot's settings.ini file.", section, key)
		if err := s.File.SaveTo(s.Source); err != nil {
			return nil, err
		}

	}
	return value, nil
}

func (s *Settings) loadBaseSection(cfg *ini.File) error {
	// Pull the base section locations
	sec, err := cfg.GetSection("")
	if err != nil {
		return err
	}

	// If we are using a supported keyring backend, then we don't need to set
	// a pin.
	for _, backend := range keyring.AvailableBackends() {
		//nolint:exhaustive
		switch backend {
		case keyring.KeychainBackend:
			return nil
		case keyring.WinCredBackend:
			return nil
		case keyring.SecretServiceBackend:
			return nil
		}
	}
	return s.getPin(sec)
}

func (s *Settings) getPin(sec *ini.Section) error {
	// Set or prompt for the sso_username variable
	pin, err := sec.GetKey("file_pin")
	if err != nil {
		pinInt := promptInt("Please enter a keychain 6 digit pin")
		pin, err = sec.NewKey("file_pin", strconv.Itoa(pinInt))
		if err != nil {
			return err
		}
		pin.Comment = "pin used to unlock file-based keyrings"
		s.pin = pinInt
		return nil
	}
	s.pin = pin.MustInt()
	return nil
}

// prompt is a simple helper function to prompt for missing
// config data.
func prompt(text string) string {
	p := promptui.Prompt{Label: text}
	resp, err := p.Run()
	if err != nil {
		logrus.Fatalf("prompt(%s)", err)
	}
	return strings.TrimSpace(strings.ToLower(resp))
}

// promptInt is a simple helper function to prompt for missing
// config data.
func promptInt(text string) (pin int) {
	resp := prompt(text)
	pin, err := strconv.Atoi(resp)
	if err != nil {
		return promptInt(text)
	}
	return pin
}

// promptSignature is function to pass into the keyring package
// that matches the signature required for the FilePrompt, but providing the
// credential's back into the keyring package.
func promptSignature(_ string) (string, error) {
	cfg, err := GetConfig("")
	if err != nil {
		return "", err
	}
	return strconv.Itoa(cfg.pin), nil
}
