package config

import (
	"encoding/json"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

var testFilename = "/tmp/config.json"

func InitTestConfig(obj any) error {
	data, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	if err := os.WriteFile(testFilename, data,
		filesystem.PermRegular); err != nil {

		return err
	}

	if err := LoadConfig(testFilename); err != nil {
		return err
	}

	return nil
}
