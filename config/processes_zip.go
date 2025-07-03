package config

import (
	"slices"
	"strings"

	"github.com/ZXSQ1/rviv/info"
)

type ZipProcessValidation struct{}

func (meta ZipProcessValidation) Name() string {
	return "zip"
}

func (meta ZipProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"archive": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"field 'archive' is not found "+
							"or has invalid format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				archiveRaw := val.(string)
				archive := NewPath(archiveRaw)
				devnames := []string{}

				LoadDevices()

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, archive.Devname) {
					info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						archive.Filename, archive.Devname, parent,
					)
				}
			},

			func(val any, parent Field) {
				archive := val.(string)

				if !strings.HasSuffix(archive, ".zip") {
					info.Error(
						"path '%s' must end in the 'zip' extension "+
							"(it is a zip archive) in field '%s'", archive, parent,
					)
				}
			},
		},

		"safe": {
			func(val any, parent Field) {
				if _, ok := val.(bool); !ok {
					info.Error(
						"field 'safe' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},
		},

		"necessary": {
			func(val any, parent Field) {
				if _, ok := val.(bool); !ok {
					info.Error(
						"field 'necessary' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},
		},
	}
}
