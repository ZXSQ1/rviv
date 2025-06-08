package sftpfs

import "github.com/ZXSQ1/rviv/logging"

// closes the filesystem (i.e. SFTP connection)
func (client *SFtpFs) Close() error {
	err := client.conn.Close()
	logging.ReportErr(err)

	return err
}
