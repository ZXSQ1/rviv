package ftpfs

import (
	"bytes"
	"io"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/jlaffaye/ftp"
)

type File struct {
	conn     *ftp.ServerConn
	filename string
	mode     filesystem.OpenMode
	wpos     uint64
	rpos     uint64
}

func (file *File) Write(p []byte) (n int, err error) {
	if file.mode != filesystem.ModeWrite {
		if file.mode == filesystem.ModeClosed {
			return -1, filesystem.ErrModeClosed
		}

		return -1, filesystem.ErrModeWrite
	}

	err = stderr(
		file.conn.StorFrom(file.filename, bytes.NewReader(p), file.wpos),
	)

	if err != nil {
		return -1, err
	}

	file.wpos += uint64(len(p))

	return len(p), nil
}

func (file *File) Read(p []byte) (n int, err error) {
	if file.mode != filesystem.ModeRead {
		if file.mode == filesystem.ModeClosed {
			return -1, filesystem.ErrModeClosed
		}

		return -1, filesystem.ErrModeRead
	}

	resp, err := file.conn.RetrFrom(
		file.filename, file.rpos,
	)

	err = stderr(err)

	if err != nil {
		return -1, err
	}

	defer resp.Close()
	n, err = resp.Read(p)
	err = stderr(err)

	if err != nil {
		if err.Error() == io.EOF.Error() {
			return 0, err
		}

		return -1, err
	}

	file.rpos += uint64(n)

	return n, nil
}

func (file *File) Close() error {
	if file.mode == filesystem.ModeClosed {
		return filesystem.ErrModeClosed
	}

	file.conn = nil
	file.filename = ""
	file.mode = filesystem.ModeClosed
	file.wpos = 0
	file.rpos = 0

	return nil
}
