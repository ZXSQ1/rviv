package webdavfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *WebDavFs) Create(filename string) error {
	if client.IsExist(filename) {
		return filesystem.ErrExist
	}

	return client.conn.Write(filename, []byte{}, filesystem.PermRegular)
}

func (client *WebDavFs) CreateDir(filename string) error {
	if client.IsExist(filename) {
		return filesystem.ErrExist
	}

	return client.conn.Mkdir(filename, filesystem.PermDir)
}
