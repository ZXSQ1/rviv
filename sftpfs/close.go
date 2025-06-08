package sftpfs

import "github.com/ZXSQ1/rviv/logging"

func (client *SFtpFs) Close() error {
	err := client.conn.Close()
	logging.ReportErr(err)

	return err
}
