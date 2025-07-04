package config

import (
	"github.com/ZXSQ1/rviv/info"
	"github.com/spf13/viper"
)

func LoadConfig(filename string) error {
	viper.SetConfigFile(filename)

	if viper.ReadInConfig() != nil {
		return info.Error(
			"could not load configuration from file '%s'", filename,
		)
	}

	return nil
}
