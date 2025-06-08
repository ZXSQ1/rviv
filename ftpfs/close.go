package ftpfs

import "github.com/ZXSQ1/rviv/logging"

// closes the filesystem (i.e. the connection)
func (client *FtpFs) Close() error {
	err := client.conn.Quit()
	logging.ReportErr(err)

	return err
}
