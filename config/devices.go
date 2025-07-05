package config

import (
	"strconv"

	"github.com/spf13/viper"
)

func LoadDevices() ([]Device, error) {
	devices := []Device{}

	for field, validations := range MainDevicesValidation.Validations() {
		val := viper.Get(string(field))

		for _, validation := range validations {
			if err := validation(val, ""); err != nil {
				return nil, err
			}
		}
	}

	devicesRaw := viper.Get("devices").(map[string]any)
	validationMap := map[string]FieldValidationMap{}

	for _, verification := range DeviceValidations {
		validationMap[verification.Name()] = verification.Validations()
	}

	for devname := range devicesRaw {
		prefix := "devices." + devname + "."
		kind := viper.GetString(prefix + "type")

		device := Device{}
		deviceInfo := DeviceInfo{}

		for field := range validationMap[kind] {
			if field == "port" {
				port, _ := strconv.Atoi(StdPath(
					viper.GetString(prefix + string(field)),
				))

				deviceInfo.Port = port
				continue
			}

			value, _ := viper.Get(prefix + string(field)).(string)

			switch field {
			case "ip":
				deviceInfo.Ip = StdPath(value)
			case "user":
				deviceInfo.User = StdPath(value)
			case "pass":
				deviceInfo.Pass = StdPath(value)
			case "prefix":
				deviceInfo.Prefix = StdPath(value)
			}
		}

		device.Name = devname
		device.Kind = kind
		device.Info = deviceInfo

		devices = append(devices, device)
	}

	return devices, nil
}
