package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type MkdirProcessValidation struct{}

func (meta MkdirProcessValidation) Name() string {
	return "mkdir"
}

func (meta MkdirProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"paths": {
			func(val any, parent Field) error {
				if paths, ok := val.([]any); !ok {
					return info.Error(
						"field 'paths' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				} else {
					for _, pathname := range paths {
						if _, ok := pathname.(string); !ok {
							return info.Error(
								"field 'paths' has invalid path '%s' "+
									"in field '%s'", pathname, parent,
							)
						}
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				pathsRaw := val.([]any)
				paths := []Path{}
				devnames := []string{}

				devices, err := LoadDevices()

				if err != nil {
					return err
				}

				for _, dev := range devices {
					devnames = append(devnames, dev.Name)
				}

				for _, pathRaw := range pathsRaw {
					pathObj, err := NewPath(pathRaw.(string))

					if err != nil {
						return err
					}

					paths = append(paths, pathObj)
				}

				for _, path := range paths {
					devname := path.Devname

					if !slices.Contains(devnames, devname) {
						return info.Error(
							"path '%s' has unknown device '%s' in field '%s'",
							path.Filename, path.Devname, parent,
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

		"parent": {
			func(val any, parent Field) error {
				if _, ok := val.(bool); !ok {
					return info.Error(
						"field 'parent' is not found "+
							"or has invalid format in field '%s'", parent,
					)
				}

				return nil
			},
		},
	}
}
