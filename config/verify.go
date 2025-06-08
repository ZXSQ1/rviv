package config

import (
	"strconv"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/spf13/viper"
)

func Verify(filename string) error {
	err := LoadConfig(filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	err = VerifyDevices()
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	err = VerifyProcesses()
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	return nil
}

func VerifyProcesses() error {
	err := IsKeyExist("processes")
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	for procIdx := range len(viper.Get("processes").([]any)) {
		procPrefix := "processes." + strconv.Itoa(procIdx) + "."

		err = IsKeyExist(procPrefix+"type", procPrefix+"name", procPrefix+"aliases")
		logging.ReportErr(err)

		if err != nil {
			return err
		}
	}

	return nil
}

func VerifyDevices() error {
	err := IsKeyExist("devices")
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	for deviceIdx := range len(viper.Get("devices").([]any)) {
		devicePrefix := "devices." + strconv.Itoa(deviceIdx) + "."

		err = IsKeyExist(devicePrefix + "type")
		logging.ReportErr(err)

		if err != nil {
			return err
		}
	}

	return nil
}
