package sftpfs

import (
	"io"

	"github.com/ZXSQ1/rviv/logging"
)

// opens a file giving a ReadWriteCloser, given a path
func (client *SFtpFs) Open(filename string) (io.ReadWriteCloser, error) {
	fileObj, err := client.conn.Open(filename)
	logging.ReportErr(err)

	return fileObj, err
}
