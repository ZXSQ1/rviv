package ftpfs

import (
	"io/fs"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *FtpFs) Stat(filename string) (fs.FileInfo, error) {
	if !client.IsExist(filename) {
		return nil, filesystem.ErrNotExist
	}

	return &FileInfo{
		filename: filename,
		conn:     client.conn,
	}, nil
}
