package ftpfs

import (
	"bytes"
	"io"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/jlaffaye/ftp"
)

// the structure implementing the io.ReadWriteCloser acting as a file object to
// be returned in the Open method
type File struct {
	// the server connection
	conn *ftp.ServerConn

	// the path to the file
	filename string

	// the position from which to write
	wpos uint64

	// the position from which to read
	rpos uint64
}

// writes the given bytes to the file returning the number of written bytes and
// an error
func (file *File) Write(p []byte) (n int, err error) {
	err = file.conn.StorFrom(
		file.filename, bytes.NewReader(p), file.wpos,
	)

	logging.ReportErr(err)

	if err != nil {
		return 0, err
	}

	file.wpos += uint64(len(p))

	return len(p), nil
}

// reads to the given buffer returning the number of read bytes and an error
func (file *File) Read(p []byte) (n int, err error) {
	resp, err := file.conn.RetrFrom(
		file.filename, file.rpos,
	)

	logging.ReportErr(err)

	if err != nil {
		return 0, err
	}

	defer resp.Close()
	n, err = resp.Read(p)

	logging.ReportErr(err)

	if err != nil {
		if err == io.EOF {
			return 0, err
		}

		return 0, err
	}

	file.rpos += uint64(n)

	return n, nil
}

// closes the file object
func (file *File) Close() error {
	file.conn = nil
	file.filename = ""
	file.wpos = 0
	file.rpos = 0

	return nil
}
