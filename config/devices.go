package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	Devices = []Device{}
)

func LoadDevices() error {
	if len(Devices) > 0 {
		return nil
	}

	for field, validations := range MainDevicesValidation.Validations() {
		for _, validation := range validations {
			if err := validation(viper.Get(string(field)), ""); err != nil {
				return err
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

		device.Name = devname
		device.Kind = kind
		device.Info = deviceInfo

		Devices = append(Devices, device)
	}

	return nil
}
