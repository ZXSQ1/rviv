package localfs

import (
	"io"
	"os"
	"strings"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func (local *LocalFs) Open(filename string, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	filename = strings.TrimLeft(filename, "/")

	switch mode {
	case filesystem.ModeRead:
		return local.fsys.OpenFile(filename, os.O_RDONLY, filesystem.PermRegular)
	case filesystem.ModeWrite:
		return local.fsys.OpenFile(filename, os.O_WRONLY, filesystem.PermRegular)
	default:
		return nil, filesystem.ErrModeInvalid
	}
}
