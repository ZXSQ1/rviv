package config

import (
	"os"
	"testing"
)

func TestFtpDevice(t *testing.T) {
	t.Cleanup(func() {
		os.Remove(testFilename)
	})

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type": "ftp",
					"ip":   "1.2.3.4",
					"port": 133,
					"user": "$USER",
					"pass": "$PASS",
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
					"type": "ftp",
					"ip":   "1.2.3.4",
					"port": 133333,
					"user": "$USER",
					"pass": "$PASS",
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
				"dev": map[string]any{
					"type": "ftp",
					"ip":   "lan",
					"port": 133,
					"user": "$USER",
					"pass": "$PASS",
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
					"type": "ftp",
					"ip":   "1.2.3.4",
					"port": -1,
					"user": "$USER",
					"pass": "$PASS",
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
				"dev": map[string]any{
					"type": "ftp",
					"ip":   "1.2.3.4",
					"port": 133,
					"user": "$USER",
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
				"dev": map[string]any{
					"type": "ftp",
					"port": 133,
					"user": "$USER",
					"pass": "$PASS",
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
				"dev": map[string]any{
					"type": "ftp",
					"ip":   "1.2.3.4",
					"port": 133,
					"pass": "$PASS",
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
				"dev": map[string]any{
					"type": "ftp",
					"ip":   "1.2.3.4",
					"user": "$USER",
					"pass": "$PASS",
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
