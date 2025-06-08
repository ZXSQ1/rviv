package sftpfs

import "github.com/ZXSQ1/rviv/logging"

func (client *SFtpFs) Create(filename string) error {
	fileObj, err := client.conn.Create(filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer fileObj.Close()
	fileObj.Write([]byte(""))

	return nil
}

func (client *SFtpFs) CreateDir(filename string) error {
	err := client.conn.MkdirAll(filename)
	logging.ReportErr(err)

	return err
}
