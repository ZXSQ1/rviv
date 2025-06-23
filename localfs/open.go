package localfs

import (
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFs) Open(filename string, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	switch mode {
	case filesystem.ModeRead:
		return os.OpenFile(filename, os.O_RDONLY, filesystem.PermRegular)
	case filesystem.ModeWrite:
		return os.OpenFile(filename, os.O_WRONLY, filesystem.PermRegular)
	default:
		return nil, filesystem.ErrModeInvalid
	}
}
