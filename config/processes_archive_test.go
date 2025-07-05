package config

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/info"
)

func TestArchiveProcess(t *testing.T) {
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
							"type":    "archive",
							"archive": "de2v:://fjskdjfl.tar",
							"entries": []string{
								"dev::/fsdfs",
							},

							"expirydays":  123,
							"compression": "xz",
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   true,
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
							"safe":        false,
							"necessary":   true,
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
							"safe":        312312,
							"necessary":   true,
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
							"safe":        true,
							"necessary":   312312,
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
							"safe":        true,
							"necessary":   true,
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
