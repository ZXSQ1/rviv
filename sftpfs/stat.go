package sftpfs

import (
	"io/fs"
)

func (client *SFtpFs) Stat(filename string) (fs.FileInfo, error) {
	return client.conn.Stat(filename)
}
