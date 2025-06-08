package sftpfs

import "github.com/ZXSQ1/rviv/logging"

// creates a file given a path
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

// creates a directory given a path
func (client *SFtpFs) CreateDir(filename string) error {
	err := client.conn.Mkdir(filename)
	logging.ReportErr(err)

	return err
}
