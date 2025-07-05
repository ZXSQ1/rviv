package config

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/info"
)

func TestMkdirProcess(t *testing.T) {
	t.Cleanup(func() {
		info.ExitOnError = true
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

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "mkdir",
							"paths": []string{
								"dev2::/dsajd",
							},

							"parent":    true,
							"necessary": true,
						},
					},
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadProcesses(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type":      "mkdir",
							"paths":     []string{},
							"parent":    true,
							"necessary": true,
						},
					},
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadProcesses(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "mkdir",
							"paths": []string{
								"dev::/dsajd",
							},

							"parent":    21903,
							"necessary": true,
						},
					},
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadProcesses(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "mkdir",
							"paths": []string{
								"dev::/dsajd",
							},

							"parent":    true,
							"necessary": 12390,
						},
					},
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadProcesses(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "mkdir",
							"paths": []string{
								"dev::/dsajd",
							},

							"parent": true,
						},
					},
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadProcesses(); err == nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"devices": map[string]any{
				"dev": map[string]any{
					"type":   "local",
					"prefix": "/",
				},
			},

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "mkdir",
							"paths": []string{
								"dev::/dsajd",
							},

							"parent":    true,
							"necessary": true,
						},
					},
				},
			},
		},
	) != nil {
		t.FailNow()
	}

	if _, err := LoadProcesses(); err != nil {
		t.FailNow()
	}
}
