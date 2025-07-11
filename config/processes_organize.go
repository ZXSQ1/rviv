package config

import (
	"regexp"
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type OrganizeProcessValidation struct{}

func (meta OrganizeProcessValidation) Name() string {
	return "organize"
}

func (meta OrganizeProcessValidation) Validations() FieldValidationMap {
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
					source, err := NewPath(sourceRaw)

					if err != nil {
						return err
					}

					sources = append(
						sources, source,
					)
				}

				devices, err := LoadDevices()

				if err != nil {
					return err
				}

				for _, dev := range devices {
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

			func(val any, parent Field) error {
				srcs := val.([]any)

				if len(srcs) == 0 {
					return info.Error(
						"field 'srcs' has no entries in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"organizedir": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'organizedir' is not found or has "+
							"invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				organizeDirRaw := val.(string)
				organizeDir, err := NewPath(organizeDirRaw)
				devnames := []string{}

				if err != nil {
					return err
				}

				devices, err := LoadDevices()

				if err != nil {
					return err
				}

				for _, dev := range devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, organizeDir.Devname) {
					return info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						organizeDir.Filename, organizeDir.Devname, parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				organizeDir := val.(string)

				if organizeDir == "" {
					return info.Error(
						"organize directory path is empty in field '%s'", parent,
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
				case "alpha", "ext", "date":
					return nil
				}

				return info.Error(
					"field 'method' has unknown type in field '%s'", parent,
				)
			},
		},

		"date": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'date' is unknown or has "+
							"invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				datestr := val.(string)
				pattern := `%[a-zA-Z]`
				re := regexp.MustCompile(pattern)

				if !re.MatchString(datestr) {
					return info.Error(
						"field 'date' does not contain any datetime specifier "+
							"in field '%s'", parent,
					)
				}

				return nil
			},
		},
	}
}
