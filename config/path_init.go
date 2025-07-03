package config

import (
	"os"
	"strings"

	"github.com/ZXSQ1/rviv/info"
)

func NewPath(uri string) Path {
	if !strings.Contains(uri, "::") {
		info.Error("must specify device in path '%s'", uri)
	}

	pathParts := strings.Split(uri, "::")
	pathDev := pathParts[0]
	pathFilename := os.ExpandEnv(pathParts[1])
	devExists := false

	for _, device := range Devices {
		if device.Name == pathDev {
			devExists = true
		}
	}

	if !devExists {
		info.Error("device '%s' not found", pathDev)
	}

	return Path{
		Filename: pathFilename,
		Devname:  pathDev,
		Active:   false,
	}
}
