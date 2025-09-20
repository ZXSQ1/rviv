package config

import (
	"os"
	"testing"
)

func TestLocalDevice(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
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
				"dev": map[string]any{
					"type": "local",
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadDevices(); err == nil {
		t.FailNow()
	}
}
