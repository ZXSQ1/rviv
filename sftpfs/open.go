package sftpfs

import "io"

func (client *SFtpFs) Open(filename string) (io.ReadWriteCloser, error) {
	return client.conn.Open(filename)
}
