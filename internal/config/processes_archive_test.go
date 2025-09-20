package config

import (
	"os"
	"testing"
)

func TestArchiveProcess(t *testing.T) {
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

			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.zip",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries":     []string{},
							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries":     []string{},
							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "fjksdf",
							"compression": true,
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1y1M1d1h1m1s",
							"compression": true,
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": 2189,
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "gz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "bz2",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": 123123,
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d21m",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0.3,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       9,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       90,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       -1,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       "fsdfsd",
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents":     []string{},
							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "dev::/fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},

							"expiry":      "1d",
							"compression": "xz",
							"level":       0,
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
							"type":        "archive",
							"archivename": "fjskdjfl.tar",
							"parents": []string{
								"dev::/fsdfs",
							},

							"entries": []string{
								"dev::/fsdfs",
							},
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
