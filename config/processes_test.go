package config

import (
	"os"
	"strings"
	"testing"

	"github.com/ZXSQ1/rviv/filesystem"
)

func TestProcesses(t *testing.T) {
	testFilename := "/tmp/config.json"

	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if os.WriteFile(testFilename, []byte(strings.TrimSpace(testConfig)),
		filesystem.PermRegular) != nil {

		t.FailNow()
	}

	LoadConfig(testFilename)
	LoadProcesses()

	if Processes == nil {
		t.FailNow()
	}
}
