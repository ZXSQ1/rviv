package config

import (
	"strconv"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/ftpfs"
	"github.com/ZXSQ1/rviv/info"
	"github.com/ZXSQ1/rviv/lan"
)

func ConnectFtpFs(infoRaw any, verbose bool) (filesystem.Filesystem, error) {
	ftpInfo, ok := infoRaw.(FtpFsConfig)

	if !ok {
		return nil, info.Error("unable to load ftp filesystem info")
	}

	var ips = []string{}
	var fsys filesystem.Filesystem

	if ftpInfo.Ip == "lan" {
		ips = lan.ScanLanAddrs(int(ftpInfo.Port), filesystem.DefaultTimeout)
	} else {
		ips = append(ips, ftpInfo.Ip)
	}

	for _, ip := range ips {
		client, err := ftpfs.Connect(&filesystem.ConnInfo{
			Addr:    ip + ":" + strconv.Itoa(int(ftpInfo.Port)),
			User:    ftpInfo.User,
			Pass:    ftpInfo.Pass,
			Timeout: filesystem.DefaultTimeout,
		})

		if err != nil {
			info.Text(
				verbose, "failed connecting '%s'", ip,
			)

			continue
		}

		info.Text(
			verbose, "succeeded connecting '%s'", ip,
		)

		fsys = client
	}

	if fsys == nil {
		return nil, info.Error("unable to connect to ftp filesystem")
	}

	return fsys, nil
}
