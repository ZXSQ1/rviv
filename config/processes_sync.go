package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type SyncProcessValidation struct{}

func (meta SyncProcessValidation) Name() string {
	return "sync"
}

func (meta SyncProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"src": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"field 'src' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				srcRaw := val.(string)
				src := NewPath(srcRaw)
				devnames := []string{}

				LoadDevices()

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, src.Devname) {
					info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						src.Filename, src.Devname, parent,
					)
				}
			},
		},

		"dest": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"field 'dest' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				destRaw := val.(string)
				dest := NewPath(destRaw)
				devnames := []string{}

				LoadDevices()

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, dest.Devname) {
					info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						dest.Filename, dest.Devname, parent,
					)
				}
			},
		},

		"oneway": {
			func(val any, parent Field) {
				if _, ok := val.(bool); !ok {
					info.Error(
						"field 'oneway' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},
		},

		"method": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"field 'method' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				method := val.(string)

				switch method {
				case "add", "remove":
					return
				}

				info.Error(
					"field 'method' has unknown type '%s' in field '%s'",
					method, parent,
				)
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
