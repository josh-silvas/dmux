package keyring

import (
	"fmt"
	"log/slog"
	"os/user"
	"strconv"
	"time"

	"github.com/99designs/keyring"
	"github.com/josh-silvas/nbot/nlog"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

var (
	userProfile *user.User
	l           nlog.Logger
)

type (
	// GetOpts : allowed settings for the Get() function in the keyring.
	GetOpts struct {
		user             string
		field            string
		promptUsr        bool
		expire           time.Duration
		promptExp        bool
		userPromptText   string
		passwdPromptText string
	}

	// Option defines a signature for a basic option passed into the client.
	Option func(c *GetOpts)

	// Logger defines a signature type that should be used to pass in any
	// logger types into the package.
	Logger func(v ...interface{})
)

// WithUser : Passes a username into the get function.
func WithUser(u string) Option {
	return func(o *GetOpts) {
		o.user = u
	}
}

// WithUserPromptText : Optionally pass in a custom prompt text for the user.
func WithUserPromptText(text string) Option {
	return func(o *GetOpts) {
		o.userPromptText = text
	}
}

// WithPasswdPromptText : Optionally pass in a custom prompt text for the password.
func WithPasswdPromptText(text string) Option {
	return func(o *GetOpts) {
		o.passwdPromptText = text
	}
}

// WithExpire : Passes an expiration into the get function.
func WithExpire(e time.Duration) Option {
	return func(o *GetOpts) {
		o.expire = e
	}
}

// PromptExpire : Prompts for an expiration into the get function.
func PromptExpire() Option {
	return func(o *GetOpts) {
		o.promptExp = true
	}
}

// PromptUser : Determines that the Get function has a username not just an API key.
func PromptUser() Option {
	return func(o *GetOpts) {
		o.promptUsr = true
	}
}

// New function will initialize a logger type, gather profile information
// and set up the config directory if needed.
func New(logger nlog.Logger) (s Settings, err error) {
	userProfile, err = user.Current()
	if err != nil {
		return s, err
	}
	if err = CreateIfNotExist(userProfile.HomeDir); err != nil {
		return s, err
	}
	cfg, err := GetConfig("")
	if err != nil {
		return s, err
	}

	// If we are at debug level in logrus, set debug in keyring
	keyring.Debug = logger.Level() == slog.LevelDebug
	cfg.Key = make(map[string]keyring.Keyring)

	cfg.Key["default"], err = open("default")
	if err != nil {
		logrus.Errorf("keyring:open:default:%s", err)
	}

	// Check if the new keychain is unlocked. If not process the unlock command.
	return cfg, keychainUnlock(cfg)
}

// APIKey : Helper method on the Get method used to get an API key or service key with no username.
func (s *Settings) APIKey(service string) (key Credential, err error) {
	return s.Get(service)
}

// UserPass : Helper method on the Get method used to get a user and password by the service name
func (s *Settings) UserPass(service string) (key Credential, err error) {
	return s.Get(service, PromptUser())
}

// UserPassCustom : Helper method on the Get method used to get a user that is custom stored by
// the calling plugin.
func (s *Settings) UserPassCustom(username string) (key Credential, err error) {
	return s.Get(username, WithUser(username), PromptExpire())
}

// Get : Method on the setting to fetch a credential. Takes a service name, such as
// `netbox` or `radius`. Pass in the `hasUser=true` if the credential stores a username
// as well as a password/api key.
func (s *Settings) Get(service string, opts ...Option) (key Credential, err error) {
	var (
		// Rewrite the service name to the full path, if not already set.
		svc = svcName(service)

		// Sets the default settings for fetching/creating a key in the keyring.
		set = &GetOpts{
			expire:           defaultExpire * 36,
			userPromptText:   fmt.Sprintf("Enter `%s` username", svc),
			passwdPromptText: fmt.Sprintf("Enter `%s` password", svc),
		}

		// Sets the credential expiration timeout.
		exp = time.Unix(time.Now().Unix(), 0).Add(set.expire).Unix()
	)

	// Override the default settings with any passed into the client.
	for _, o := range opts {
		o(set)
	}

	if svc == "default" {
		return s.SetCredential(s.User, svc, svc, exp)
	}

	if key, err = s.GetCredential(s.User, svc); err == nil {
		key.Username, key.Keys[KeyName] = splitter(key.Keys[KeyName])
		return key, nil
	}

	// USERNAME PROMPT : Used to set a username value if it's not passed in.
	if set.promptUsr {
		pUser := promptui.Prompt{Label: set.userPromptText}
		set.user, err = pUser.Run()
		if err != nil {
			logrus.Fatalf("%s.PromptUser(%s)", svc, err)
		}
	}

	// PASSWORD GATHER : Logic used to fetch a password.
	pPass := promptui.Prompt{Label: set.passwdPromptText, Mask: '*'}
	pass, err := pPass.Run()
	if err != nil {
		logrus.Fatalf("%s.PromptPass(%s)", svc, err)
	}

	// Small function to hash the user/password if a user exists, or return the password.
	password := func() string {
		if set.user != "" {
			return fmt.Sprintf("%s%s%s", set.user, Split, pass)
		}
		return pass
	}()

	// EXPIRE PROMPT : If the flag is passed in, then prompt for a custom expire timeout.
	if set.promptExp {
		pExp := promptui.Prompt{Label: fmt.Sprintf("Enter `%s` password expiration in days", svc), Default: "365"}
		rExp, err := pExp.Run()
		if err != nil {
			logrus.Fatalf("%s.PromptExpire(%s)", svc, err)
		}
		i, err := strconv.Atoi(rExp)
		if err != nil {
			logrus.Fatalf("%s.PromptExpire(%s)", svc, err)
		}
		exp = time.Unix(time.Now().Unix(), 0).Add(time.Hour * 24 * time.Duration(i)).Unix()
	}

	return s.SetCredential(s.User, password, svc, exp)
}

// Delete function will clean any associated keyring for a given user.
func (s *Settings) Delete(service string) error {
	svc := svcName(service)
	if !s.keyExists(svc) {
		var err error
		s.Key[svc], err = open(svc)
		if err != nil {
			return fmt.Errorf("service.open:%w", err)
		}
	}

	if err := s.Key[svc].Remove(svcUser(s.User, svc)); err != nil {
		return fmt.Errorf("service.Remove:%w", err)
	}

	logrus.Infof("deleted `%s` key", svc)
	return nil
}
