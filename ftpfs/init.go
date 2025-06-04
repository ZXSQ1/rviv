package ftpfs

import (
	"strconv"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/jlaffaye/ftp"
)

type FtpFs struct {
	conn *ftp.ServerConn
}

func Connect(ip string, port int, user, pass string) (filesystem.Filesystem, error) {
	portString := strconv.Itoa(port)
	addr := ip + ":" + portString
	conn, err := ftp.Dial(addr)

	if err != nil {
		return nil, err
	}

	if conn.Login(user, pass) != nil {
		return nil, err
	}

	return &FtpFs{conn}, nil
}
