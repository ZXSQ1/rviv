package ftpfs

import "github.com/ZXSQ1/rviv/logging"

func (client *FtpFs) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return err.Error() != "no such file or directory"
	}

	return true
}
