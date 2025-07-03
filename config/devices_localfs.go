package config

import (
	"github.com/ZXSQ1/rviv/info"
)

type LocalDeviceValidation struct{}

func (meta LocalDeviceValidation) Name() string {
	return "local"
}

func (meta LocalDeviceValidation) Validations() FieldValidationMap {
	return map[Field][]Validation{
		"prefix": {
			func(val any, parent Field) {
				switch val.(type) {
				case string:
					return
				default:
					info.Error(
						"key 'prefix' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},
		},
	}
}
