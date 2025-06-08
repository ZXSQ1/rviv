package ftpfs

import "github.com/ZXSQ1/rviv/logging"

// removes the file given a path
func (client *FtpFs) Remove(filename string) error {
	err := client.conn.Delete(filename)
	logging.ReportErr(err)

	return err
}

// removes the directory given a path
func (client *FtpFs) RemoveDir(filename string) error {
	err := client.conn.RemoveDirRecur(filename)
	logging.ReportErr(err)

	return err
}
