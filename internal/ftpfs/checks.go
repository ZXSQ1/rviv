package ftpfs

import (
	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/ZXSQ1/rviv/internal/logging"
)

func (client *FtpFs) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return err.Error() != filesystem.ErrNotExist.Error()
	}

	return true
}
