package ftpfs

import (
	"bytes"
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

// creates a file with the given path; returns an os.ErrExist error if the file
// already exists (default handling of this in FTP library is not found)
func (client *FtpFs) Create(filename string) error {
	if client.IsExist(filename) {
		logging.ReportErr(os.ErrExist)
		return os.ErrExist
	}

	err := client.conn.Stor(filename, bytes.NewReader([]byte{}))
	logging.ReportErr(err)

	return err
}

// creates a directory with the given path; returns an os.ErrExist err if the
// directory already exists (by default, the FTP libraru does not count
// attempting to create an existent directory an error)
func (client *FtpFs) CreateDir(filename string) error {
	if client.IsExist(filename) {
		logging.ReportErr(os.ErrExist)
		return os.ErrExist
	}

	err := client.conn.MakeDir(filename)
	logging.ReportErr(err)

	return err
}
