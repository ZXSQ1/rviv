package ftpfs

import "github.com/ZXSQ1/rviv/logging"

func (client *FtpFs) Close() error {
	err := client.conn.Quit()
	logging.ReportErr(err)

	return err
}
