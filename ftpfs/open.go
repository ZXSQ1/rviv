package ftpfs

import (
	"io"
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (client *FtpFs) Open(filename string) (io.ReadWriteCloser, error) {
	if !client.IsExist(filename) {
		logging.ReportErr(os.ErrNotExist)
		return nil, os.ErrNotExist
	}

	stat, err := client.Stat(filename)
	logging.ReportErr(err)

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
