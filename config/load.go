package config

import (
	"github.com/ZXSQ1/rviv/logging"
	"github.com/spf13/viper"
)

func LoadConfig(filename string) error {
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	return nil
}
