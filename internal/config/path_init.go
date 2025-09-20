package config

import (
	"slices"
	"strings"

	"github.com/ZXSQ1/rviv/internal/info"
)

var DeviceSeparator = "::"

func NewPath(uri string) (Path, error) {
	if !strings.Contains(uri, DeviceSeparator) {
		return Path{}, info.Error("must specify device in path '%s'", uri)
	}

	pathParts := strings.Split(uri, DeviceSeparator)
	pathDev := pathParts[0]
	pathFilename := StdPath(pathParts[1])
	devnames := []string{}
	devices, err := LoadDevices()

	if err != nil {
		return Path{}, info.Error("unable to load devices")
	}

	for _, device := range devices {
		devnames = append(devnames, device.Name)
	}

	if !slices.Contains(devnames, pathDev) {
		return Path{}, info.Error("device '%s' not found", pathDev)
	}

	return Path{
		Devices:  devices,
		Filename: pathFilename,
		Devname:  pathDev,
		Active:   false,
	}, nil
}
