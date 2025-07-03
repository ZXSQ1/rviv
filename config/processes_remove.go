package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type RemoveProcessValidation struct{}

func (meta RemoveProcessValidation) Name() string {
	return "remove"
}

func (meta RemoveProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"paths": {
			func(val any, parent Field) {
				if _, ok := val.([]string); !ok {
					info.Error(
						"field 'paths' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				pathsRaw := val.([]string)
				paths := []Path{}
				devnames := []string{}

				LoadDevices()

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				for _, pathRaw := range pathsRaw {
					paths = append(paths, NewPath(pathRaw))
				}

				for _, path := range paths {
					devname := path.Devname

					if !slices.Contains(devnames, devname) {
						info.Error(
							"path '%s' has unknown device '%s' in field '%s'",
							path.Filename, path.Devname, parent,
						)
					}
				}
			},
		},

		"recursive": {
			func(val any, parent Field) {
				if _, ok := val.(bool); !ok {
					info.Error(
						"field 'recursive' is not found or "+
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
