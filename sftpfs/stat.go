package sftpfs

import (
	"io/fs"
)

func (client *SFtpFs) Stat(filename string) (fs.FileInfo, error) {
	stat, err := client.conn.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		conn:     client.conn,
		stat:     stat,
		filename: filename,
	}, nil
}
