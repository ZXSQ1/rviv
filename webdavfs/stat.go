package webdavfs

import (
	"io/fs"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *WebDavFs) Stat(filename string) (fs.FileInfo, error) {
	if !client.IsExist(filename) {
		return nil, filesystem.ErrNotExist
	}

	return &FileInfo{
		filename: filename,
		conn:     client.conn,
	}, nil
}
