package config

import (
	"fmt"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/spf13/viper"
)

// checks if given keys are found in the configuration; assumes that the viper
// configuration is loaded (using a function like LoadConfig); the keys are in
// the format specified by viper in the Get function
func IsKeyExist(keys ...string) error {
	for _, key := range keys {
		if viper.Get(key) == nil {
			err := fmt.Errorf("the '%s' field is not found", key)
			logging.ReportErr(err)

			return err
		}
	}

	return nil
}
