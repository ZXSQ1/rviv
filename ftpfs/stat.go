package ftpfs

import (
	"io/fs"
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

// gives information about the file given a path; equivalent to the os.Stat
// function for FTP; returns an implementation of the FileInfo interface (i.e.
// the ftpfs.FileInfo structure)
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
