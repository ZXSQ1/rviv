package ftpfs

import (
	"io"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *FtpFs) Open(filename string, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	if !client.IsExist(filename) {
		return nil, filesystem.ErrNotExist
	}

	stat, err := client.Stat(filename)
	err = stderr(err)

	if err != nil {
		return nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, filesystem.ErrFileNotRegular
	}

	if !(mode == filesystem.ModeRead || mode == filesystem.ModeWrite) {
		return nil, filesystem.ErrModeInvalid
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		mode:     mode,
	}, nil
}
