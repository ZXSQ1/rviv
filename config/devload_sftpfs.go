package config

func LoadSFtpFsConfig(deviceRaw map[string]any) any {
	return SFtpFsConfig{
		Ip:   StdPath(deviceRaw["ip"].(string)),
		Port: uint16(deviceRaw["port"].(float64)),
		User: StdPath(deviceRaw["user"].(string)),
		Pass: StdPath(deviceRaw["pass"].(string)),
	}
}
