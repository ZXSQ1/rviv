package sftpfs

import "github.com/ZXSQ1/rviv/logging"

// removes a file given a path
func (client *SFtpFs) Remove(filename string) error {
	err := client.conn.Remove(filename)
	logging.ReportErr(err)

	return err
}

// removes a directory given a path
func (client *SFtpFs) RemoveDir(filename string) error {
	err := client.conn.RemoveAll(filename)
	logging.ReportErr(err)

	return err
}
