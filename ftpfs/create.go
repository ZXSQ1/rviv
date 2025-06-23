package ftpfs

import (
	"bytes"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *FtpFs) Create(filename string) error {
	if client.IsExist(filename) {
		return filesystem.ErrExist
	}

	return stderr(
		client.conn.Stor(filename, bytes.NewReader([]byte{})),
	)
}

func (client *FtpFs) CreateDir(filename string) error {
	if client.IsExist(filename) {
		return filesystem.ErrExist
	}

	return stderr(client.conn.MakeDir(filename))
}
