package keyring

import (
	"errors"
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/99designs/keyring"
)

// KeyName default name of the key stored in keychain
const KeyName = "__key__"

// Credential type is used as the type set/retrieved when
// interacting with the keyring package
type Credential struct {
	Username string
	Keys     map[string]string
	Expire   int64
}

// GetCredential function will take a username and service type to get a user credential. The
// expiry timer is required here to determine if the password returned from the keyring service
// is actually valid.
func (s *Settings) GetCredential(user string, service string) (cred Credential, err error) {
	// Check to see if the service-name exist in the Settings.Key map. If not
	// create the entry in the map and open the keyring.
	if !s.keyExists(service) {
		s.Key[service], err = open(service)
		if err != nil {
			return cred, fmt.Errorf("service.open:%w", err)
		}
	}

	// Fetch the credential from the keyring.
	key, err := s.Key[service].Get(svcUser(user, service))
	if err != nil {
		return cred, fmt.Errorf("service.Get:%w", err)
	}
	parsed, err := parseCredential(key)
	if err != nil {
		return cred, fmt.Errorf("parseCredential:%w", err)
	}

	if !parsed.isExpired() {
		parsed.Username = user
		return parsed, nil
	}

	return parsed, errors.New("password is expired")
}

// SetCredential function is a small wrapper to the keyring Set function, but with the
// formatting of the password to match this packages Expire types.
func (s *Settings) SetCredential(user, key, service string, exp int64) (Credential, error) {
	service = svcName(service)

	cred := Credential{
		Username: func() string {
			u, _ := splitter(key)
			if u != "" {
				return u
			}
			return user
		}(),
		Keys:   map[string]string{KeyName: key},
		Expire: exp,
	}
	var item = keyring.Item{
		Key:         svcUser(user, service),
		Data:        []byte(fmt.Sprintf("%d  %s", cred.Expire, key)),
		Label:       cases.Title(language.English).String(svcShort(service)),
		Description: service,
	}

	if !s.keyExists(service) {
		var err error
		s.Key[service], err = open(service)
		if err != nil {
			return Credential{}, fmt.Errorf("unable to open service `%s`", service)
		}
	}

	if err := s.Key[service].Set(item); err != nil {
		l.Errorf("error setting credential SetCredential:keyring.Set: %s", err)
		return cred, fmt.Errorf("SetCredential.Set:%w", err)
	}
	return cred, nil
}
