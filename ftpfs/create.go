package ftpfs

import (
	"bytes"
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (client *FtpFs) Create(filename string) error {
	if client.IsExist(filename) {
		logging.ReportErr(os.ErrExist)
		return os.ErrExist
	}

	err := client.conn.Stor(filename, bytes.NewReader([]byte{}))
	logging.ReportErr(err)

	return err
}

func (client *FtpFs) CreateDir(filename string) error {
	if client.IsExist(filename) {
		logging.ReportErr(os.ErrExist)
		return os.ErrExist
	}

	err := client.conn.MakeDir(filename)
	logging.ReportErr(err)

	return err
}
