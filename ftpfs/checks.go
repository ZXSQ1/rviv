package ftpfs

import (
	"github.com/jlaffaye/ftp"
)

func (client *Client) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	return status(err) != ftp.StatusFileUnavailable
}
