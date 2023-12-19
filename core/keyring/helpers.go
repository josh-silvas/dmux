package keyring

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/keyring"
)

// open : Opens up a keyring using a standard configuration.
func open(s string) (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{
		AllowedBackends: []keyring.BackendType{
			keyring.KeychainBackend,
			keyring.WinCredBackend,
			keyring.FileBackend,
		},
		ServiceName: svcName(s),

		// Needed for default file fallback
		FileDir:          fileBackend,
		FilePasswordFunc: promptSignature,

		// MacOS default items
		KeychainName:                   KeyChainName,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: true,
	})
}

func keychainUnlock(cfg Settings) error {
	if !isMacOS() {
		return nil
	}

	if !cfg.Test {
		// Attempt a pull for SSO to see if the keychain exist, or if we need to create a new one.
		if _, err := cfg.Get("default"); err != nil {
			return fmt.Errorf("keychainUnlock.cfg.SSO:%w", err)
		}
	}

	keychainDB := fmt.Sprintf("%s.keychain-db", KeyChainName)

	//nolint:gosec
	out, err := exec.Command(keychainCMD, "show-keychain-info", keychainDB).CombinedOutput()
	if err != nil {
		return fmt.Errorf("keychainUnlock.exec.Command:show-keychain-info failed %s:%w", keychainDB, err)
	}

	// If there is no-timeout set on the keychain itself, this is what we want. Exit here.
	if strings.Contains(fmt.Sprintf("%s", out), "no-timeout") {
		return nil
	}

	//nolint:gosec
	if _, err = exec.Command(keychainCMD, "set-keychain-settings", keychainDB).CombinedOutput(); err != nil {
		return fmt.Errorf("keychainUnlock.exec.Command:set-keychain-settings failed %s:%w", keychainDB, err)
	}
	return nil
}

// logPrint function uses the Logger method associated with the non exported value.
func logPrint(v ...interface{}) {
	if logger == nil {
		return
	}
	logger(v...)
}

// splitter : Helper function to split the password field into the username and password
// if it is a hashed type.
func splitter(s string) (string, string) {
	if !strings.Contains(s, Split) {
		return "", s
	}
	sp := strings.Split(s, Split)
	return sp[0], sp[1]
}

// svcUser function will return the username used for the
// service we create for each token.
// We have to modify the user for backends that do not support a
// service name in their storage
func svcUser(user, service string) string {
	if isLinux() {
		return fmt.Sprintf("%s:%s", service, user)
	}
	return user
}

// svcName : returns the full name of the service.
func svcName(s string) string {
	if strings.HasPrefix(s, SvcBase) {
		return s
	}
	return fmt.Sprintf("%s.%s", SvcBase, s)
}

// svcShort : returns the short name of the service.
func svcShort(s string) string {
	if strings.HasPrefix(s, SvcBase) {
		return strings.Replace(s, SvcBase+".", "", -1)
	}
	return s
}

// isExpired method will take a Expire time with a Credential receiver and
// determine if the Credential time is past the passed in Expire.
func (c *Credential) isExpired() bool {
	return c.Expire < time.Now().Unix()
}

// keyExists : Returns true if the key exist in the local map of the opened
// keyrings. NOTE: This will not contain ALL keys, just ones that have been
// opened.
func (s *Settings) keyExists(svc string) bool {
	if _, ok := s.Key[svcName(svc)]; !ok {
		return false
	}
	return true
}

// parseCredential function will take a string from the keyring service that has both the
// user password and the Expire time and return a Credential structure with the parsed values.
func parseCredential(s keyring.Item) (Credential, error) {
	var resp Credential
	parsed := strings.Split(string(s.Data), "  ")
	if len(parsed) != 2 {
		return resp, fmt.Errorf("unable to parse secret, got len(%d) %v", len(parsed), parsed)
	}
	exp, err := strconv.ParseInt(parsed[0], 10, 64)
	if err != nil {
		logPrint("error at parseCredential/strconv.ParseInt")
		return resp, err
	}
	return Credential{
		Expire: exp,
		Keys:   map[string]string{KeyName: parsed[1]},
	}, nil
}

// Password : Helper method on the Credential struct to retrieve the default key name.
func (c *Credential) Password() string {
	if strings.Contains(c.Keys[KeyName], Split) {
		_, pass := splitter(c.Keys[KeyName])
		return pass
	}
	return c.Keys[KeyName]
}

// isMacOS : Returns true if the platform is a MacOS type.
func isMacOS() bool {
	return runtime.GOOS == "darwin"
}

// isLinux : Returns true if the platform is one of the supported Linux types.
func isLinux() bool {
	for _, platform := range []string{"linux", "freebsd", "linux2"} {
		if platform == runtime.GOOS {
			return true
		}
	}
	return false
}
