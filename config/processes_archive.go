package config

import (
	"path/filepath"
	"slices"
	"strings"

	"github.com/ZXSQ1/rviv/expiry"
	"github.com/ZXSQ1/rviv/info"
)

type ArchiveProcessValidation struct{}

func (meta ArchiveProcessValidation) Name() string {
	return "archive"
}

func (meta ArchiveProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"parents": {
			func(val any, parent Field) error {
				if parents, ok := val.([]any); !ok {
					return info.Error(
						"field 'parents' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				} else {
					for _, entry := range parents {
						if _, ok := entry.(string); !ok {
							return info.Error(
								"field 'parents' has invalid path '%s' "+
									"in field '%s'", entry, parent,
							)
						}
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				parentsRaw := val.([]any)
				parents := []Path{}
				devnames := []string{}

				devices, err := LoadDevices()

				if err != nil {
					return err
				}

				for _, dev := range devices {
					devnames = append(devnames, dev.Name)
				}

				for _, parentRaw := range parentsRaw {
					parentEntry, err := NewPath(parentRaw.(string))

					if err != nil {
						return err
					}

					parents = append(parents, parentEntry)
				}

				for _, parentEntry := range parents {
					devname := parentEntry.Devname

					if !slices.Contains(devnames, devname) {
						return info.Error(
							"path '%s' has unknown device '%s' in field '%s'",
							parentEntry.Filename, parentEntry.Devname, parent,
						)
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				parents := val.([]any)

				if len(parents) == 0 {
					return info.Error(
						"field 'parents' has no entries in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"archivename": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'archivename' is not found "+
							"or has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				archivename := val.(string)

				if archivename == "" {
					return info.Error(
						"archive name is empty in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				archivename := val.(string)
				ext := filepath.Ext(archivename)

				if ext == ".tar" || ext == ".zip" {
					return nil
				}

				return info.Error(
					"path '%s' must end in 'tar' or 'zip' extensions "+
						"in field '%s'", archivename, parent,
				)
			},

			func(val any, parent Field) error {
				archivename := val.(string)

				if !strings.Contains(archivename, string(filepath.Separator)) {
					return nil
				}

				return info.Error(
					"path '%s' must have one component only (basename "+
						"of an archive) in field '%s'", archivename, parent,
				)
			},
		},

		"entries": {
			func(val any, parent Field) error {
				if entries, ok := val.([]any); !ok {
					return info.Error(
						"field 'entries' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				} else {
					for _, entry := range entries {
						if _, ok := entry.(string); !ok {
							return info.Error(
								"field 'entries' has invalid path '%s' "+
									"in field '%s'", entry, parent,
							)
						}
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				entriesRaw := val.([]any)
				entries := []Path{}
				devnames := []string{}

				devices, err := LoadDevices()

				if err != nil {
					return err
				}

				for _, dev := range devices {
					devnames = append(devnames, dev.Name)
				}

				for _, entryRaw := range entriesRaw {
					entry, err := NewPath(entryRaw.(string))

					if err != nil {
						return err
					}

					entries = append(entries, entry)
				}

				for _, entry := range entries {
					devname := entry.Devname

					if !slices.Contains(devnames, devname) {
						return info.Error(
							"path '%s' has unknown device '%s' in field '%s'",
							entry.Filename, entry.Devname, parent,
						)
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				entries := val.([]any)

				if len(entries) == 0 {
					return info.Error(
						"field 'entries' has no entries in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"expiry": {
			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'expiry' has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				expiryfmt := val.(string)

				if _, err := expiry.ParseExpiry(expiryfmt); err != nil {
					return err
				}

				return nil
			},
		},

		"compression": {
			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'compression' has "+
							"invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				compressionMethod := val.(string)

				switch compressionMethod {
				case "xz", "gz", "bz2", "":
					return nil
				}

				return info.Error(
					"field 'compression' has unknown compression "+
						"method '%s' in field '%s'", compressionMethod, parent,
				)
			},
		},

		"level": {
			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				if _, ok := val.(float64); !ok {
					return info.Error(
						"field 'level' has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				level := val.(float64)

				if float64(int(level)) != level {
					return info.Error(
						"field 'level' is not an integer in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				if val == nil {
					return nil
				}

				level := val.(float64)

				if level < 0 || level > 9 {
					return info.Error(
						"field 'level' must be betweeen 0 and 9 in field '%s'",
						parent,
					)
				}

				return nil
			},
		},
	}
}
