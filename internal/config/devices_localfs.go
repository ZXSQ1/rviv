package config

import (
	"github.com/ZXSQ1/rviv/internal/info"
)

type LocalDeviceValidation struct{}

func (meta LocalDeviceValidation) Name() string {
	return "local"
}

func (meta LocalDeviceValidation) Validations() FieldValidationMap {
	return map[Field][]Validation{
		"prefix": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"key 'prefix' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}

				return nil
			},
		},
	}
}
