package keyring

import (
	"testing"
)

func TestListKeys(t *testing.T) {
	// Clean up any existing test keys first
	for _, svc := range []string{"test1", "test2", "test3"} {
		_ = testCfg.Delete(svc)
	}

	// Set up some test keys
	testServices := []string{"test1", "test2", "test3"}
	for _, svc := range testServices {
		_, err := testCfg.SetCredential(
			testCred.Username,
			testCred.Password(),
			svcName(svc),
			testCred.Expire,
		)
		if err != nil {
			t.Fatalf("failed to set test credential for %s: %v", svc, err)
		}
	}

	// List the keys
	keys, err := testCfg.ListKeys()
	if err != nil {
		t.Fatalf("ListKeys failed: %v", err)
	}

	// Verify we got back our test services
	if len(keys) == 0 {
		t.Error("no keys were returned")
	}

	for _, svc := range testServices {
		found := false
		for _, key := range keys {
			if key == svc {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected to find service %s in keys list, got keys: %v", svc, keys)
		}
	}

	// Clean up test keys
	for _, svc := range testServices {
		if err := testCfg.Delete(svc); err != nil {
			t.Logf("warning: failed to clean up test key %s: %v", svc, err)
		}
	}
}
