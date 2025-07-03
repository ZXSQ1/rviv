package config

import (
	"github.com/ZXSQ1/rviv/info"
	"github.com/spf13/viper"
)

func LoadConfig(filename string) {
	viper.SetConfigFile(filename)

	if viper.ReadInConfig() != nil {
		info.Error("could not load configuration from file '%s'", filename)
	}
}
