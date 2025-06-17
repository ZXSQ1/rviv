package webdavfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *WebDavFs) Create(filename string) error {
	return client.conn.Write(filename, []byte{}, filesystem.PermRegular)
}

func (client *WebDavFs) CreateDir(filename string) error {
	return client.conn.Mkdir(filename, filesystem.PermDir)
}
