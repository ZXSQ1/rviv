package config

func LoadWebDavFsConfig(deviceRaw map[string]any) any {
	return WebDavFsConfig{
		Ip:   StdPath(deviceRaw["ip"].(string)),
		Port: uint16(deviceRaw["port"].(float64)),
		User: StdPath(deviceRaw["user"].(string)),
		Pass: StdPath(deviceRaw["pass"].(string)),
	}
}
