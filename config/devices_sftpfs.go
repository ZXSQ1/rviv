package config

import (
	"net"

	"github.com/ZXSQ1/rviv/info"
)

type SFtpDeviceValidation struct{}

func (meta SFtpDeviceValidation) Name() string {
	return "ssh"
}

func (meta SFtpDeviceValidation) Validations() FieldValidationMap {
	return map[Field][]Validation{
		"ip": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"key 'ip' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}

				return nil
			},

			func(val any, parent Field) error {
				if ip := val.(string); !(net.ParseIP(ip) != nil || ip == "lan") {
					return info.Error(
						"key 'ip' has bad format in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"port": {
			func(val any, parent Field) error {
				if port, ok := val.(float64); ok {
					return nil
				} else if float64(int(port)) == port {
					return nil
				}

				return info.Error(
					"key 'port' has wrong type or "+
						"is not found in field '%s'", parent,
				)
			},

			func(val any, parent Field) error {
				if port, _ := val.(float64); port > 65535 || port < 0 {
					return info.Error(
						"key 'port' has bad value (too high "+
							"or too low) '%s'", parent,
					)
				}

				return nil
			},
		},

		"user": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"key 'user' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}

				return nil
			},
		},

		"pass": {
			func(val any, parent Field) error {
				if _, ok := val.(string); !ok {
					return info.Error(
						"key 'pass' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}

				return nil
			},
		},
	}
}
