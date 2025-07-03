package config

import (
	"regexp"

	"github.com/ZXSQ1/rviv/info"
)

type FtpDeviceValidation struct{}

func (meta FtpDeviceValidation) Name() string {
	return "ftp"
}

func (meta FtpDeviceValidation) Validations() FieldValidationMap {
	return map[Field][]Validation{
		"ip": {
			func(val any, parent Field) {
				switch val.(type) {
				case string:
					return
				default:
					info.Error(
						"key 'ip' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				pattern := `\d.\d.\d.\d`
				re := regexp.MustCompile(pattern)
				valString, _ := val.(string)

				if !re.Match([]byte(valString)) || valString != "lan" {
					info.Error(
						"key 'ip' has bad format in field '%s'", parent,
					)
				}
			},
		},

		"port": {
			func(val any, parent Field) {
				switch val.(type) {
				case int:
					return
				default:
					info.Error(
						"key 'port' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				valInt, _ := val.(int)

				if valInt > 65535 || valInt < 0 {
					info.Error(
						"key 'port' has bad value (too high "+
							"or too low) '%s'", parent,
					)
				}
			},
		},

		"user": {
			func(val any, parent Field) {
				switch val.(type) {
				case string:
					return
				default:
					info.Error(
						"key 'user' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},
		},

		"pass": {
			func(val any, parent Field) {
				switch val.(type) {
				case string:
					return
				default:
					info.Error(
						"key 'pass' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},
		},
	}
}
