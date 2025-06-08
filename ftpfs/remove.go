package ftpfs

import "github.com/ZXSQ1/rviv/logging"

func (client *FtpFs) Remove(filename string) error {
	err := client.conn.Delete(filename)
	logging.ReportErr(err)

	return err
}

func (client *FtpFs) RemoveDir(filename string) error {
	err := client.conn.RemoveDirRecur(filename)
	logging.ReportErr(err)

	return err
}
