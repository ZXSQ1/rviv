package config

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/info"
)

func TestMainDevice(t *testing.T) {
	t.Cleanup(func() {
		info.ExitOnError = true
		os.Remove(testFilename)
	})

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev12": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadDevices(); err != nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev12": map[string]any{
					"type":   "locl",
					"prefix": "/",
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadDevices(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev#": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadDevices(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev_-12": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadDevices(); err != nil {
		t.FailNow()
	}
}
