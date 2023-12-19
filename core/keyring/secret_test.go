package keyring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setCredential(t *testing.T) {
	c, err := testCfg.SetCredential(testCred.Username, testCred.Password(), "com.keyring.nbot.test", testCred.Expire)
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, c.Password(), testCred.Password()) {
		t.Fatalf("ERROR: assert.Equal %s", c.Password())
	}
}

func Test_getCredential(t *testing.T) {
	c, err := testCfg.GetCredential(testCred.Username, "com.keyring.nbot.test")
	if err != nil {
		t.Fatal(err)
	}
	if !assert.Equal(t, c.Password(), testCred.Password()) {
		t.Fatalf("ERROR: assert.Equal %s", c.Password())
	}
}

func Test_isExpired(t *testing.T) {
	if expired := testCred.isExpired(); expired {
		t.Fatal("ERROR: incorrect Expire")
	}
	testCred.Expire = testCred.Expire - 1200
	if expired := testCred.isExpired(); !expired {
		t.Fatal("ERROR: incorrect Expire")
	}
}
