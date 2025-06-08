package ftpfs

import (
	"io/fs"
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (client *FtpFs) Stat(filename string) (fs.FileInfo, error) {
	if !client.IsExist(filename) {
		logging.ReportErr(os.ErrNotExist)
		return nil, os.ErrNotExist
	}

	return &FileInfo{
		filename: filename,
		conn:     client.conn,
	}, nil
}
