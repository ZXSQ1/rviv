package sftpfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

// checks if the file exists given a path
func (client *SFtpFs) IsExist(filename string) bool {
	_, err := client.conn.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
