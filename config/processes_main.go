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
			func(val any, parent Field) error {
				if _, ok := val.(map[string]any); !ok {
					return info.Error(
						"field 'processes' is not found or has invalid format",
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				processesRaw := val.(map[string]any)

				for groupname, groupinfo := range processesRaw {
					groupinfo := groupinfo.(map[string]any)

					for _, c := range groupname {
						if ('A' > c || 'Z' < c) && ('a' > c || 'z' < c) &&
							!(c == '_' || c == '-') && ('0' > c || '9' < c) {

							return info.Error(
								"process group '%s' has invalid characters "+
									"(A-Z, a-z, _ and - are only allowed) in field "+
									"'processes'", groupname,
							)
						}
					}

					if aliases, ok := groupinfo["aliases"].([]any); !ok {
						return info.Error(
							"field 'aliases' is not found or has invalid "+
								"format in field 'processes.%s'", groupname,
						)
					} else {
						for _, alias := range aliases {
							if _, ok := alias.(string); !ok {
								return info.Error(
									"field 'aliases' has invalid alias '%s' "+
										"in field 'processes.%s'", alias, groupname,
								)
							}

							alias := alias.(string)

							for _, c := range alias {
								if ('A' > c || 'Z' < c) && ('a' > c || 'z' < c) &&
									!(c == '_' || c == '-') && ('0' > c || '9' < c) {

									return info.Error(
										"alias '%s' has invalid characters "+
											"(A-Z, a-z, _ and - are only allowed) in "+
											"field 'processes.%s'", alias, groupname,
									)
								}
							}
						}
					}

					if subprocesses, ok := groupinfo["subprocesses"].([]any); !ok {
						return info.Error(
							"field 'subprocesses' is not found or has invalid "+
								"format in field 'processes.%s'", groupname,
						)
					} else {
						for _, subprocess := range subprocesses {
							if _, ok := subprocess.(map[string]any); !ok {
								return info.Error(
									"field 'subprocesses' has invalid subprocess "+
										"'%s' in field 'processes.%s'", subprocess,
									groupname,
								)
							}
						}
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				processesRaw := val.(map[string]any)
				validationMap := map[string]FieldValidationMap{}
				prockinds := []string{}

				for _, validation := range ProcessVerifications {
					validationMap[validation.Name()] = validation.Validations()
					prockinds = append(prockinds, validation.Name())
				}

				for groupname, groupinfo := range processesRaw {
					groupinfo := groupinfo.(map[string]any)
					prefix := "processes." + groupname + ".subprocesses"
					subprocesses := groupinfo["subprocesses"].([]any)

					for _, subprocess := range subprocesses {
						subprocess := subprocess.(map[string]any)

						if _, ok := subprocess["type"].(string); !ok {
							return info.Error(
								"field 'type' is not found or has invalid format "+
									"in field '%s'", prefix,
							)
						}

						kind := subprocess["type"].(string)

						if !slices.Contains(prockinds, kind) {
							return info.Error(
								"subprocess field '%s' has unknown type '%s'",
								prefix, kind,
							)
						}

						for fieldName, validations := range validationMap[kind] {
							field := subprocess[string(fieldName)]

							for _, validation := range validations {
								if err := validation(field, Field(prefix)); err != nil {
									return err
								}
							}
						}
					}
				}

				return nil
			},
		},
	}
}
