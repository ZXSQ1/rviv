package config

import (
	"strconv"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/spf13/viper"
)

// verifies the existence of certain configuration keys and their types; calls
// other functions that verify the main keys (keys in the first layer of the
// configuration); not complete yet
func Verify(filename string) error {
	err := VerifyDevices()
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

// verifies the "processes" main key and its subkeys' types; not complete yet (
// other keys must be checked for existence, and key types must be verified)
func VerifyProcesses() error {
	err := IsKeyExist("processes")
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	for procIdx := range len(viper.Get("processes").([]any)) {
		procPrefix := "processes." + strconv.Itoa(procIdx) + "."

		err = IsKeyExist(
			procPrefix+"type", procPrefix+"name", procPrefix+"aliases",
		)

		logging.ReportErr(err)

		if err != nil {
			return err
		}
	}

	return nil
}

// verifies the "devices" main key and its subkeys' types; not complete yet (
// other keys must be checked for existence, and key types must be verified
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
