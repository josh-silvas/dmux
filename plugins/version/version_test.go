package version

import (
	"strings"
	"testing"
	"time"
)

var (
	ver, _     = SemVer("v1.2.0")
	testCfgVer = ConfigVersion{
		Version:   ver,
		Timestamp: time.Now(),
	}
)

func TestCfgVer_String(t *testing.T) {
	verString := testCfgVer.String()
	if !strings.Contains(verString, testCfgVer.Version.String()) {
		t.Errorf("ToString Error %s", verString)
	}
	verStruct, err := ParseConfigVersion(verString)
	if err != nil {
		t.Fatal(err)
	}
	if verStruct.Version.String() != testCfgVer.Version.String() {
		t.Errorf("Parsing error %s != %s", verStruct.Version, testCfgVer.Version)
	}
}
