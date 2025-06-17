package sftpfs

import "github.com/ZXSQ1/rviv/logging"

func (client *SFtpFs) Remove(filename string) error {
	err := client.conn.Remove(filename)
	logging.ReportErr(err)

	return err
}

func (client *SFtpFs) RemoveDir(filename string) error {
	err := client.conn.RemoveAll(filename)
	logging.ReportErr(err)

	return err
}
