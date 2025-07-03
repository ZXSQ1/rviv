package config

import (
	"slices"

	"github.com/ZXSQ1/rviv/info"
	"github.com/spf13/viper"
)

type DeviceValidation struct{}

func (meta DeviceValidation) Name() string {
	return ""
}

func (meta DeviceValidation) Validations() FieldValidationMap {
	return FieldValidationMap{
		"devices": {
			func(val any, parent Field) {
				if _, ok := val.(map[string]any); !ok {
					info.Error(
						"field 'devices' has invalid format or is not found",
					)
				}

				devicesRaw := val.(map[string]any)

				for devname, devInfo := range devicesRaw {
					if _, ok := devInfo.(map[string]any); !ok {
						info.Error(
							"field '%s' is not found or has "+
								"invalid format in field '%s'", devname, parent,
						)
					}

					for _, c := range devname {
						if ('A' > c || 'Z' < c) && ('a' > c || 'z' < c) &&
							!(c == '_' || c == '-') && ('0' > c || '9' < c) {

							info.Error(
								"device name '%s' has invalid characters "+
									"(A-Z, a-z, _ and - are only allowed) in field "+
									"'devices'", devname,
							)
						}
					}
				}
			},

			func(val any, parent Field) {
				devkinds := []string{}
				devices, _ := val.(map[string]any)

				for _, deviceVerificaition := range DeviceValidations {
					devkinds = append(devkinds, deviceVerificaition.Name())
				}

				for devname, devInfo := range devices {
					devInfo := devInfo.(map[string]any)
					devkind, ok := devInfo["type"].(string)

					if !ok {
						info.Error(
							"field 'type' has invalid format (must be string) "+
								"or is not found in field '%s'",
							"devices."+devname+".type",
						)
					}

					if !slices.Contains(devkinds, devkind) {
						info.Error(
							"field 'type' has unknown device type in field '%s'",
							"devices."+devname+".type",
						)
					}
				}
			},

			func(val any, parent Field) {
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
							validation(viper.Get(string(field)), Field(prefix))
						}
					}
				}
			},
		},
	}
}
