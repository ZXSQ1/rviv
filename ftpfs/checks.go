package ftpfs

import (
	"github.com/jlaffaye/ftp"
)

func (client *FTPFS) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	return status(err) != ftp.StatusFileUnavailable
}
