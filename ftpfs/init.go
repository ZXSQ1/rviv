package ftpfs

import (
	"strconv"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
	"github.com/jlaffaye/ftp"
)

// the structure/object implementing the Filesystem interface for FTP
type FtpFs struct {

	// the FTP connection
	conn *ftp.ServerConn
}

// connects to the FTP server given the address, the username and the password
// returning the Filesystem implementation
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

// returns the address given the ip and the port
func GetAddr(ip string, port int) string {
	portString := strconv.Itoa(port)
	return ip + ":" + portString
}
