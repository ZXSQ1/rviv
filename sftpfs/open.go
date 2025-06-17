package sftpfs

import (
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *SFtpFs) Open(filename string, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	switch mode {
	case filesystem.ModeRead:
		return client.conn.OpenFile(filename, os.O_RDONLY)
	case filesystem.ModeWrite:
		return client.conn.OpenFile(filename, os.O_WRONLY)
	default:
		return nil, os.ErrInvalid
	}
}
