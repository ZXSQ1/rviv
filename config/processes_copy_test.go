package config

import (
	"os"
	"testing"

	"github.com/ZXSQ1/rviv/info"
)

func TestCopyProcess(t *testing.T) {
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
							"type": "copy",
							"srcs": []string{
								"dev::/ehfdjsk/lfjkl",
							},
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
			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "copy",
							"srcs": []string{
								"dev::/ehfdjsk/lfjkl",
							},

							"dest":      "dev::/asdas/asdas",
							"method":    "dd",
							"necessary": true,
						},
					},
				},
			},

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

	if _, err := LoadProcesses(); err != nil {
		t.FailNow()
	}

	if InitTestConfig(
		map[string]any{
			"processes": map[string]any{
				"procgroup": map[string]any{
					"aliases": []string{"pg"},
					"subprocesses": []map[string]any{
						{
							"type": "copy",
							"srcs": []string{
								"de12v::/ehfdjsk/lfjkl",
							},

							"dest":      "dev::/asdas/asdas",
							"method":    "dd",
							"necessary": true,
						},
					},
				},
			},

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
							"type":      "copy",
							"srcs":      []string{},
							"dest":      "dev::/asdas/asdas",
							"method":    "d31d",
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
							"type": "copy",
							"srcs": []string{
								"dev::/sada",
							},

							"dest":      "dev::/asdas/asdas",
							"method":    "dd",
							"necessary": 1293,
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
							"type":      "copy",
							"srcs":      []string{},
							"dest":      "dev::/asdas/asdas",
							"method":    "dd",
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
							"type": "copy",
							"srcs": []string{
								"dev::/dsajd",
							},

							"dest":      "dev::/asdas/asdas",
							"method":    "dd",
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
