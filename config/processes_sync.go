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
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'src' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				srcRaw := val.(string)
				src := NewPath(srcRaw)
				devnames := []string{}

				LoadDevices()

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, src.Devname) {
					return info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						src.Filename, src.Devname, parent,
					)
				}

				return nil
			},
		},

		"dest": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'dest' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				destRaw := val.(string)
				dest := NewPath(destRaw)
				devnames := []string{}

				if err := LoadDevices(); err != nil {
					return err
				}

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, dest.Devname) {
					return info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						dest.Filename, dest.Devname, parent,
					)
				}

				return nil
			},
		},

		"oneway": {
			func(val any, parent Field) error {
				if _, ok := val.(bool); !ok {
					return info.Error(
						"field 'oneway' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"method": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'method' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				method := val.(string)

				switch method {
				case "add", "remove":
					return nil
				}

				return info.Error(
					"field 'method' has unknown type '%s' in field '%s'",
					method, parent,
				)
			},
		},

		"necessary": {
			func(val any, parent Field) error {
				if _, ok := val.(bool); !ok {
					return info.Error(
						"field 'necessary' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}

				return nil
			},
		},
	}
}
