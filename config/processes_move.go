package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type CopyProcessValidation struct{}

func (meta CopyProcessValidation) Name() string {
	return "copy"
}

func (meta CopyProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"srcs": {
			func(val any, parent Field) {
				if _, ok := val.([]string); !ok {
					info.Error(
						"field 'srcs' not found or has invalid "+
							"format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				sourcesRaw := val.([]string)
				sources := []Path{}
				devnames := []string{}

				for _, sourceRaw := range sourcesRaw {
					sources = append(
						sources, NewPath(sourceRaw),
					)
				}

				LoadDevices()

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				for _, source := range sources {
					devname := source.Devname

					if !slices.Contains(devnames, devname) {
						info.Error(
							"path '%s' has unknown device '%s' in field '%s'",
							source.Filename, source.Devname, parent,
						)
					}
				}
			},
		},

		"dest": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"field 'dest' is not found or has "+
							"invalid format in field '%s'", val,
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

		"method": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"field 'method' is unknown or has "+
							"invalid format in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				method := val.(string)

				switch method {
				case "ff", "fd", "dd":
					return
				}

				info.Error("field 'method' has unknown type in field '%s'", parent)
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
