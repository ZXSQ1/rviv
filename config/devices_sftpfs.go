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
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"key 'ip' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},

			func(val any, parent Field) {
				if ip := val.(string); !(net.ParseIP(ip) != nil || ip == "lan") {
					info.Error(
						"key 'ip' has bad format in field '%s'", parent,
					)
				}
			},
		},

		"port": {
			func(val any, parent Field) {
				if port, ok := val.(float64); ok {
					return
				} else if float64(int(port)) == port {
					return
				}

				info.Error(
					"key 'port' has wrong type or "+
						"is not found in field '%s'", parent,
				)
			},

			func(val any, parent Field) {
				if port, _ := val.(float64); port > 65535 || port < 0 {
					info.Error(
						"key 'port' has bad value (too high "+
							"or too low) '%s'", parent,
					)
				}
			},
		},

		"user": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"key 'user' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},
		},

		"pass": {
			func(val any, parent Field) {
				if _, ok := val.(string); !ok {
					info.Error(
						"key 'pass' has wrong type or "+
							"is not found in field '%s'", parent,
					)
				}
			},
		},
	}
}
