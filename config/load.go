package config

import (
	"github.com/ZXSQ1/rviv/logging"
	"github.com/spf13/viper"
)

// loads the viper configuration given a filename; needs only to be loaded once
// (viper apparently stores the read config throughout the program), but does
// not necessarily mean that something will go wrong if it is loaded multiple
// times
func LoadConfig(filename string) error {
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	return nil
}
