package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type MoveProcessValidation struct{}

func (meta MoveProcessValidation) Name() string {
	return "copy"
}

func (meta MoveProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"srcs": {
			func(val any, parent Field) error {
				if srcs, ok := val.([]any); !ok {
					return info.Error(
						"field 'srcs' not found or has invalid "+
							"format in field '%s'", parent,
					)
				} else {
					for _, src := range srcs {
						if _, ok := src.(string); !ok {
							return info.Error(
								"field 'srcs' has invalid source '%s' "+
									"in field '%s'", src, parent,
							)
						}
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				sourcesRaw := val.([]any)
				sources := []Path{}
				devnames := []string{}

				for _, sourceRaw := range sourcesRaw {
					sourceRaw := sourceRaw.(string)
					sources = append(
						sources, NewPath(sourceRaw),
					)
				}

				if err := LoadDevices(); err != nil {
					return err
				}

				for _, dev := range Devices {
					devnames = append(devnames, dev.Name)
				}

				for _, source := range sources {
					devname := source.Devname

					if !slices.Contains(devnames, devname) {
						return info.Error(
							"path '%s' has unknown device '%s' in field '%s'",
							source.Filename, source.Devname, parent,
						)
					}
				}

				return nil
			},
		},

		"dest": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'dest' is not found or has "+
							"invalid format in field '%s'", val,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				destRaw := val.(string)
				dest := NewPath(destRaw)
				devnames := []string{}

				LoadDevices()

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

		"method": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'method' is unknown or has "+
							"invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				method := val.(string)

				switch method {
				case "ff", "fd", "dd":
					return nil
				}

				return info.Error(
					"field 'method' has unknown type in field '%s'", parent,
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
