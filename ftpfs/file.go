package ftpfs

import (
	"bytes"
	"io"

	"github.com/jlaffaye/ftp"
)

type File struct {
	conn     *ftp.ServerConn
	filename string
	wpos     uint64
	rpos     uint64
}

func (file *File) Write(p []byte) (n int, err error) {
	err = file.conn.StorFrom(
		file.filename, bytes.NewReader(p), file.wpos,
	)

	if err != nil {
		return -1, err
	}

	file.wpos += uint64(len(p))

	return len(p), nil
}

func (file *File) Read(p []byte) (n int, err error) {
	resp, err := file.conn.RetrFrom(
		file.filename, file.rpos,
	)

	if err != nil {
		return -1, err
	}

	defer resp.Close()
	n, err = resp.Read(p)

	if err != nil {
		if err == io.EOF {
			return 0, err
		}

		return -1, err
	}

	file.rpos += uint64(n)

	return n, nil
}

func (file *File) Close() error {
	file.conn = nil
	file.filename = ""
	file.wpos = 0
	file.rpos = 0

	return nil
}
