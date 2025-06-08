package sftpfs

import (
	"io/fs"

	"github.com/ZXSQ1/rviv/logging"
)

// gives information about the file given a path
func (client *SFtpFs) Stat(filename string) (fs.FileInfo, error) {
	fileObj, err := client.conn.Stat(filename)
	logging.ReportErr(err)

	return fileObj, err
}
