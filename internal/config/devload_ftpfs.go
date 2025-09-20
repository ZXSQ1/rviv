package config

func LoadFtpFsConfig(deviceRaw map[string]any) any {
	return FtpFsConfig{
		Ip:   StdPath(deviceRaw["ip"].(string)),
		Port: uint16(deviceRaw["port"].(float64)),
		User: StdPath(deviceRaw["user"].(string)),
		Pass: StdPath(deviceRaw["pass"].(string)),
	}
}
