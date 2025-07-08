package config

import (
	"slices"
	"strings"

	"github.com/ZXSQ1/rviv/info"
)

type ArchiveProcessValidation struct{}

func (meta ArchiveProcessValidation) Name() string {
	return "archive"
}

func (meta ArchiveProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"archive": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'archive' is not found "+
							"or has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				archiveRaw := val.(string)
				archive, err := NewPath(archiveRaw)

				if err != nil {
					return err
				}

				devnames := []string{}

				devices, err := LoadDevices()

				if err != nil {
					return err
				}

				for _, dev := range devices {
					devnames = append(devnames, dev.Name)
				}

				if !slices.Contains(devnames, archive.Devname) {
					return info.Error(
						"path '%s' has unknown device '%s' in field '%s'",
						archive.Filename, archive.Devname, parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				archive := val.(string)

				if strings.HasSuffix(archive, ".tar") {
					return nil
				} else if strings.HasSuffix(archive, ".zip") {
					return nil
				}

				return info.Error(
					"path '%s' must end in 'tar' or 'zip' extensions "+
						"in field '%s'", archive, parent,
				)
			},

			func(val any, parent Field) error {
				archivename := val.(string)

				if archivename == "" {
					return info.Error(
						"archive path is empty in field '%s'", parent,
					)
				}

				return nil
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

		"expirydays": {
			func(val any, parent Field) error {
				if _, ok := val.(float64); !ok {
					return info.Error(
						"field 'expirydays' is not found "+
							"or has invalid format in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"compression": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"field 'compression' is not found "+
							"or has invalid format in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
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

		"safe": {
			func(val any, parent Field) error {
				if _, ok := val.(bool); !ok {
					return info.Error(
						"field 'safe' is not found or "+
							"has invalid format in field '%s'", parent,
					)
				}

				return nil
			},
		},
	}
}
