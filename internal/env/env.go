package env

import "os"

var (
	ConfigFilename = os.Getenv("RVIV_CONFIG")
	Verbose        = true
	Colored        = true
)
