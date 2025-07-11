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
							"type":    "archive",
							"archive": "de2v:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        true,
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
							"archive":     "dev:://fjskdjfl.zip",
							"entries":     []string{},
							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        true,
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
							"archive":     "dev:://fjskdjfl.tar",
							"entries":     []string{},
							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  "fjksdf",
							"compression": true,
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": 2189,
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "",
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "gz",
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "bz2",
							"level":       0,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": 123123,
							"level":       0,
							"safe":        false,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        312312,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       0.3,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       9,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       90,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       -1,
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       "fsdfsd",
							"safe":        true,
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
							"type":    "archive",
							"archive": "dev:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"level":       0,
							"safe":        true,
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
