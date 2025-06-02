package ftpfs

import (
	"bytes"
	"os"

	"github.com/jlaffaye/ftp"
)

func (client *Client) CreateFile(filename string) error {
	exist, err := client.IsExist(filename)

	if err != nil {
		return err
	}

	if exist {
		return os.ErrExist
	}

	return client.conn.Stor(filename, bytes.NewReader([]byte{}))
}

func CreateDir(conn *ftp.ServerConn, filename string) error {
	return conn.MakeDir(filename)
}
