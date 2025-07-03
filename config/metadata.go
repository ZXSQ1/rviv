package config

type DeviceInfo struct {
	Ip     string
	Port   int
	User   string
	Pass   string
	Prefix string
}

type Device struct {
	Name string
	Kind string
	Info DeviceInfo
}

type Process struct {
	Kind    string
	Options any
}

type ProcessGroup struct {
	Groupname string
	Aliases   []string
	Processes []Process
}
