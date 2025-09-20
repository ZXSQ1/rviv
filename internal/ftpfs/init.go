package ftpfs

import (
	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/jlaffaye/ftp"
)

type FtpFs struct {
	conn     *ftp.ServerConn
	connInfo *filesystem.ConnInfo
}

func Connect(connInfo *filesystem.ConnInfo) (filesystem.Filesystem, error) {
	conn, err := ftp.Dial(connInfo.Addr,
		ftp.DialWithTimeout(connInfo.Timeout),
	)

	err = stderr(err)

	if err != nil {
		return nil, err
	}

	err = conn.Login(connInfo.User, connInfo.Pass)
	err = stderr(err)

	if err != nil {
		return nil, err
	}

	return &FtpFs{conn: conn, connInfo: connInfo}, nil
}
