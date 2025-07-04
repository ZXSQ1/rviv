package config

import (
	"os"
	"strings"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestDevices(t *testing.T) {
	testFilename := "/tmp/config.json"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.WriteFile(testFilename, []byte(strings.TrimSpace(testConfig)),
		filesystem.PermRegular) != nil {

		t.FailNow()
	}

	LoadConfig(testFilename)
	LoadDevices()

	if len(Devices) == 0 {
		t.FailNow()
	}
}
