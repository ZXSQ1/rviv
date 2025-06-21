package ftpfs

import (
	"fmt"
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *FtpFs) Open(filename string, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	if !client.IsExist(filename) {
		return nil, os.ErrNotExist
	}

	stat, err := client.Stat(filename)
	err = stderr(err)

	if err != nil {
		return nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, fmt.Errorf("open operation failed; file is not regular")
	}

	if !(mode == filesystem.ModeRead || mode == filesystem.ModeWrite) {
		return nil, os.ErrInvalid
	}

	return &File{
		conn:     client.conn,
		filename: filename,
		mode:     mode,
	}, nil
}
