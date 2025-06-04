package ftpfs

import (
	"bytes"
	"os"
)

func (client *FtpFs) Create(filename string) error {
	if client.IsExist(filename) {
		return os.ErrExist
	}

	return client.conn.Stor(filename, bytes.NewReader([]byte{}))
}

func (client *FtpFs) CreateDir(filename string) error {
	if client.IsExist(filename) {
		return os.ErrExist
	}

	return client.conn.MakeDir(filename)
}
