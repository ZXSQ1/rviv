package config

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestFtpDevice(t *testing.T) {
	testFilename := "/tmp/config.json"
	testConfig, err := json.Marshal(map[string]any{})

	if err != nil {
		t.FailNow()
	}

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.WriteFile(testFilename, testConfig, filesystem.PermRegular) != nil {
		t.FailNow()
	}
}
