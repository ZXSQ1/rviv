package config

import (
	"strconv"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/ftpfs"
	"github.com/ZXSQ1/rviv/info"
	"github.com/ZXSQ1/rviv/lan"
)

var (
	ActiveConns = map[string]filesystem.Filesystem{}
)

func (file *Path) Connect(necessary bool) error {
	if conn, ok := ActiveConns[file.Devname]; ok {
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

	var fsys filesystem.Filesystem
	var err error

	switch file.Devname {
	case "ftp", "ssh", "webdav":
		ips := []string{}

		if device.Info.Ip == "lan" {
			ips = lan.ScanLanAddrs(device.Info.Port, filesystem.DefaultTimeout)
		} else {
			ips = append(ips, device.Info.Ip)
		}

		for _, ip := range ips {
			connInfo := &filesystem.ConnInfo{
				Addr: ip + ":" + strconv.Itoa(device.Info.Port),
				User: device.Info.User,
				Pass: device.Info.Pass,
			}

			switch device.Kind {
			case "ftp":
				fsys, err = ftpfs.Connect(connInfo)

				if err != nil {
					continue
				}
			case "ssh":
				fsys, err = ftpfs.Connect(connInfo)

				if err != nil {
					continue
				}
			case "webdav":
				fsys, err = ftpfs.Connect(connInfo)

				if err != nil {
					continue
				}
			}
		}

		if necessary {
			return info.Error("unable to connect to device '%s'", device.Name)
		} else {
			return info.Warning("unable to connect to device '%s'", device.Name)
		}
	}

	ActiveConns[file.Devname] = fsys
	file.Fsys = fsys
	file.Active = true

	return nil
}
