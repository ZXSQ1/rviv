package config

func LoadLocalFsConfig(deviceRaw map[string]any) any {
	return LocalFsConfig{
		Prefix: deviceRaw["prefix"].(string),
	}
}
