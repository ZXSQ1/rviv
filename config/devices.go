package config

import (
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
		prefix := "devices." + devname
		kind := viper.GetString(prefix + ".type")
		deviceRaw := viper.Get(prefix).(map[string]any)

		device := Device{}
		device.Kind = kind
		device.Name = devname

		switch kind {
		case "local":
			device.Info = LoadLocalFsConfig(deviceRaw)
		case "ftp":
			device.Info = LoadFtpFsConfig(deviceRaw)
		case "ssh":
			device.Info = LoadSFtpFsConfig(deviceRaw)
		case "webdav":
			device.Info = LoadWebDavFsConfig(deviceRaw)
		}

		devices = append(devices, device)
	}

	return devices, nil
}
