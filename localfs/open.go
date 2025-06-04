package localfs

import (
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFS) Open(filename string) (io.ReadWriteCloser, error) {
	return os.OpenFile(filename, os.O_RDWR, filesystem.RegularPerm)
}
