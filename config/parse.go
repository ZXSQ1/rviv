package config

import (
	"encoding/json"
	"os"
)

func Parse(jsonFilename string) (any, error) {
	var data any

	bytes, err := os.ReadFile(jsonFilename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, data)
	return data, err
}
