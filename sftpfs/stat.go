package sftpfs

import (
	"io/fs"
)

func (client *SFtpFs) Stat(filename string) (fs.FileInfo, error) {
	_, err := client.conn.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		conn:     client.conn,
		filename: filename,
	}, nil
}
