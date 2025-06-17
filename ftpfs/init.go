package ftpfs

import (
	"strconv"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
	"github.com/jlaffaye/ftp"
)

type FtpFs struct {
	conn *ftp.ServerConn
}

func Connect(addr, user, pass string) (filesystem.Filesystem, error) {
	conn, err := ftp.Dial(addr)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	err = conn.Login(user, pass)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	return &FtpFs{conn: conn}, nil
}

func GetAddr(ip string, port int) string {
	portString := strconv.Itoa(port)
	return ip + ":" + portString
}
