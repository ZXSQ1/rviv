package ftpfs

import (
	"io"
	"os"
)

func (client *FtpFs) Open(filename string) (io.ReadWriteCloser, error) {
	if !client.IsExist(filename) {
		return nil, os.ErrNotExist
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		wpos:     0,
		rpos:     0,
	}, nil
}
