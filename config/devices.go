package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	Devices = []Device{}
)

func LoadDevices() {
	if len(Devices) > 0 {
		return
	}

	for field, validations := range MainDevicesValidation.Validations() {
		for _, validation := range validations {
			validation(viper.Get(string(field)), "")
		}
	}

	devicesRaw := viper.Get("devices").(map[string]map[string]any)
	validationMap := map[string]FieldValidationMap{}

	for _, verification := range DeviceValidations {
		validationMap[verification.Name()] = verification.Validations()
	}

	for key := range devicesRaw {
		prefix := "devices." + key + "."
		name := key
		kind := viper.GetString(prefix + "type")

		device := Device{}
		deviceInfo := DeviceInfo{}

		for field := range validationMap[kind] {
			if field == "port" {
				port, _ := strconv.Atoi(os.ExpandEnv(
					viper.GetString(prefix + string(field)),
				))

				deviceInfo.Port = port
				continue
			}

			value, _ := viper.Get(prefix + string(field)).(string)

			switch field {
			case "ip":
				deviceInfo.Ip = os.ExpandEnv(value)
			case "user":
				deviceInfo.User = os.ExpandEnv(value)
			case "pass":
				deviceInfo.Pass = os.ExpandEnv(value)
			case "prefix":
				deviceInfo.Prefix = os.ExpandEnv(value)
			}
		}

		device.Name = name
		device.Kind = kind
		device.Info = deviceInfo

		Devices = append(Devices, device)
	}
}
