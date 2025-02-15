package keyring

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"testing"
	"time"

	"github.com/99designs/keyring"
)

var (
	testCfg  Settings
	testCred = Credential{
		Username: "test_user",
		Keys:     map[string]string{"password": "somesecret"},
		Expire:   time.Unix(time.Now().Unix(), 0).Add(10 * time.Minute).Unix(),
	}
)

func TestMain(m *testing.M) {
	var err error
	if testCfg, err = testNew(); err != nil {
		log.Fatalf("dmux.keyring.New:%s", err)
	}
	os.Exit(m.Run())
}

func testNew() (s Settings, err error) {
	keyName := "com.keyring.dmux.test"
	userProfile, err = user.Current()
	if err != nil {
		return s, fmt.Errorf("testNew:%w", err)
	}
	if err = CreateIfNotExist(userProfile.HomeDir); err != nil {
		return s, fmt.Errorf("user.Current():%w", err)
	}
	cfg, err := GetConfig("")
	if err != nil {
		return s, fmt.Errorf("testNew:%w", err)
	}
	cfg.Test = true

	cfg.Key = make(map[string]keyring.Keyring)
	cfg.Key[keyName], err = keyring.Open(keyring.Config{
		AllowedBackends: []keyring.BackendType{
			keyring.KeychainBackend,
			keyring.WinCredBackend,
			keyring.FileBackend,
		},
		ServiceName: keyName,

		// Needed for default file fallback
		FileDir:          fileBackend,
		FilePasswordFunc: promptSignature,

		// MacOS default items
		KeychainName:                   KeyChainName,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: true,
	})
	if err != nil {
		log.Fatalf("keyring:open:%s:%s", keyName, err)
	}

	// Check if the new keychain is unlocked. If not
	// process the unlock command.
	return cfg, keychainUnlock(cfg)
}
