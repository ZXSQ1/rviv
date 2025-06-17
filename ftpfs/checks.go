package ftpfs

import (
	"github.com/ZXSQ1/rviv/logging"
)

func (client *FtpFs) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	logging.ReportErr(err)

	if err != nil {
		return stderr(err).Error() == "no such file or directory"
	}

	return true
}
