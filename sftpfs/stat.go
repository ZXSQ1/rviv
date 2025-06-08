package sftpfs

import (
	"io/fs"

	"github.com/ZXSQ1/rviv/logging"
)

func (client *SFtpFs) Stat(filename string) (fs.FileInfo, error) {
	fileObj, err := client.conn.Stat(filename)
	logging.ReportErr(err)

	return fileObj, err
}
