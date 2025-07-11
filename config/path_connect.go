package config

import (
	"fmt"
	"strconv"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/ftpfs"
	"github.com/ZXSQ1/rviv/info"
	"github.com/ZXSQ1/rviv/lan"
	"github.com/ZXSQ1/rviv/localfs"
)

var (
	activeConns = map[string]filesystem.Filesystem{}
)

func (file *Path) Connect(verbose bool) error {
	info.Heading(verbose, "connect")

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

	var fsys filesystem.Filesystem
	var err error

	switch file.Devname {
	case "local":
		fsys, err = localfs.Init(device.Info.Prefix)

		if err != nil {
			return info.Error("unable to connect to device '%s'", device.Name)
		}
	case "ftp", "ssh", "webdav":
		ips := []string{}

		if device.Info.Ip == "lan" {
			ips = lan.ScanLanAddrs(device.Info.Port, filesystem.DefaultTimeout)
		} else {
			ips = append(ips, device.Info.Ip)
		}

	loop:
		for _, ip := range ips {
			connInfo := &filesystem.ConnInfo{
				Addr: ip + ":" + strconv.Itoa(device.Info.Port),
				User: device.Info.User,
				Pass: device.Info.Pass,
			}

			info.Text(verbose, "tried address '%s:%d'... ", device.Info.Ip,
				device.Info.Port)

			switch device.Kind {
			case "ftp":
				fsys, err = ftpfs.Connect(connInfo)

				if err == nil {
					break loop
				}
			case "ssh":
				fsys, err = ftpfs.Connect(connInfo)

				if err == nil {
					break loop
				}
			case "webdav":
				fsys, err = ftpfs.Connect(connInfo)

				if err == nil {
					break loop
				}
			}

			if verbose {
				fmt.Println("failed")
			}
		}

		if fsys == nil {
			return info.Error("unable to connect to device '%s'", device.Name)
		} else if verbose {
			fmt.Println("success")
		}

	}

	activeConns[file.Devname] = fsys
	file.Fsys = fsys
	file.Active = true

	return nil
}
