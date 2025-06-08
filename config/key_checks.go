package config

import (
	"fmt"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/spf13/viper"
)

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
