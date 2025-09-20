package config

type Device struct {
	Name string
	Kind string
	Info any
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
