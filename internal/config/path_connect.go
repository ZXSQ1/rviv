package config

import (
	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/ZXSQ1/rviv/internal/info"
)

var (
	activeConns = map[string]filesystem.Filesystem{}
)

func (file *Path) Connect(verbose bool) error {
	if conn, ok := activeConns[file.Devname]; ok {
		file.Fsys = conn
		file.Active = true

		return nil
	}

	var device Device

	for _, targetDevice := range file.Devices {
		if file.Devname == targetDevice.Name {
			device = targetDevice
		}
	}

	info.Text(verbose, "connecting to device '%s'", file.Devname)

	var fsys filesystem.Filesystem
	var err error

	switch device.Kind {
	case "local":
		fsys, err = ConnectLocalFs(device.Info, verbose)
	case "ftp":
		fsys, err = ConnectFtpFs(device.Info, verbose)
	case "ssh":
		fsys, err = ConnectSFtpFs(device.Info, verbose)
	case "webdav":
		fsys, err = ConnectWebDavFs(device.Info, verbose)
	}

	if err != nil {
		return info.Error("unable to connect to device '%s'", device.Name)
	}

	activeConns[file.Devname] = fsys
	file.Fsys = fsys
	file.Active = true

	return nil
}
