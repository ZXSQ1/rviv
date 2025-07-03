package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
)

type ProcessValidation struct{}

func (meta ProcessValidation) Name() string {
	return ""
}

func (meta ProcessValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"processes": {
			func(val any, parent Field) {
				if _, ok := val.(map[string]map[string]any); !ok {
					info.Error(
						"field 'processes' is not found or has invalid format",
					)
				}
			},

			func(val any, parent Field) {
				processesRaw := val.(map[string]map[string]any)

				for groupname, groupinfo := range processesRaw {
					for _, c := range groupname {
						if ('A' > c && 'Z' < c) || ('a' > c && 'z' < c) ||
							!(c == '_' || c == '-') || ('0' > c && '9' < c) {

							info.Error(
								"process group '%s' has invalid characters "+
									"(A-Z, a-z, _ and - are only allowed) in field "+
									"'processes'", groupname,
							)
						}
					}

					if aliases, ok := groupinfo["aliases"].([]string); !ok {
						info.Error(
							"field 'aliases' is not found or has invalid "+
								"format in field 'processes.%s'", groupname,
						)
					} else {
						for _, alias := range aliases {
							for _, c := range alias {
								if ('A' > c && 'Z' < c) || ('a' > c && 'z' < c) ||
									!(c == '_' || c == '-') || ('0' > c && '9' < c) {

									info.Error(
										"alias '%s' has invalid characters "+
											"(A-Z, a-z, _ and - are only allowed) in "+
											"field 'processes.%s'", alias, groupname,
									)
								}
							}
						}
					}

					if _, ok := groupinfo["subprocesses"].([]map[string]any); !ok {
						info.Error(
							"field 'subprocesses' is not found or has invalid "+
								"format in field 'processes.%s'", groupname,
						)
					}
				}
			},

			func(val any, parent Field) {
				processesRaw := val.(map[string]map[string]any)
				validationMap := map[string]FieldValidationMap{}
				prockinds := []string{}

				for _, validation := range ProcessVerifications {
					validationMap[validation.Name()] = validation.Validations()
					prockinds = append(prockinds, validation.Name())
				}

				for groupname, groupinfo := range processesRaw {
					prefix := "processes." + groupname
					subprocesses := groupinfo["subprocesses"].([]map[string]any)

					for _, subprocess := range subprocesses {
						prefix = prefix + ".subprocesses"

						if _, ok := subprocess["type"].(string); !ok {
							info.Error(
								"field 'type' is not found or has invalid format "+
									"in field '%s'", prefix,
							)
						}

						kind := subprocess["type"].(string)

						if !slices.Contains(prockinds, kind) {
							info.Error(
								"subprocess field '%s' has unknown type '%s'",
								prefix, kind,
							)
						}

						for fieldName, validations := range validationMap[kind] {
							field := subprocess[string(fieldName)]

							for _, validation := range validations {
								validation(field, Field(prefix))
							}
						}
					}
				}
			},
		},
	}
}
