package ftpfs

import (
	"io"
	"os"
)

func (client *FtpFs) Open(filename string) (io.ReadWriteCloser, error) {
	if !client.IsExist(filename) {
		return nil, os.ErrNotExist
	}

	stat, err := client.Stat(filename)

	if err != nil {
		return nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, os.ErrInvalid
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		wpos:     0,
		rpos:     0,
	}, nil
}
