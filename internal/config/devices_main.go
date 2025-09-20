package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/internal/info"
	"github.com/spf13/viper"
)

type DeviceValidation struct{}

func (meta DeviceValidation) Name() string {
	return ""
}

func (meta DeviceValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"devices": {
			func(val any, parent Field) error {
				if _, ok := val.(map[string]any); !ok {
					return info.Error(
						"field 'devices' has invalid format or is not found",
					)
				}

				devicesRaw := val.(map[string]any)

				for devname, devInfo := range devicesRaw {
					if _, ok := devInfo.(map[string]any); !ok {
						return info.Error(
							"field '%s' is not found or has "+
								"invalid format in field '%s'", devname, parent,
						)
					}

					for _, c := range devname {
						if ('A' > c || 'Z' < c) && ('a' > c || 'z' < c) &&
							!(c == '_' || c == '-') && ('0' > c || '9' < c) {

							return info.Error(
								"device name '%s' has invalid characters "+
									"(A-Z, a-z, _ and - are only allowed) in field "+
									"'devices'", devname,
							)
						}
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				devkinds := []string{}
				devices, _ := val.(map[string]any)

				for _, deviceVerificaition := range DeviceValidations {
					devkinds = append(devkinds, deviceVerificaition.Name())
				}

				for devname, devInfo := range devices {
					devInfo := devInfo.(map[string]any)
					devkind, ok := devInfo["type"].(string)

					if !ok {
						return info.Error(
							"field 'type' has invalid format (must be string) "+
								"or is not found in field '%s'",
							"devices."+devname+".type",
						)
					}

					if !slices.Contains(devkinds, devkind) {
						return info.Error(
							"field 'type' has unknown device type in field '%s'",
							"devices."+devname+".type",
						)
					}
				}

				return nil
			},

			func(val any, parent Field) error {
				devices, _ := val.(map[string]any)
				validationMap := map[string]FieldValidationMap{}

				for _, verification := range DeviceValidations {
					validationMap[verification.Name()] = verification.Validations()
				}

				for devname, devInfo := range devices {
					devInfo := devInfo.(map[string]any)
					prefix := "devices." + devname
					devkind := devInfo["type"].(string)

					for field, validations := range validationMap[devkind] {
						field = Field(prefix+".") + field

						for _, validation := range validations {
							if err := validation(
								viper.Get(string(field)), Field(prefix),
							); err != nil {
								return err
							}
						}
					}
				}

				return nil
			},
		},
	}
}
