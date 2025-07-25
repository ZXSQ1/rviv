package sftpfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *SFtpFs) CreateFile(filename string) error {
	if client.IsExist(filename) {
		return filesystem.ErrExist
	}

	fileObj, err := client.conn.Create(filename)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

func (client *SFtpFs) CreateDir(filename string) error {
	return client.conn.Mkdir(filename)
}
