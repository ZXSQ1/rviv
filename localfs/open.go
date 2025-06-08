package localfs

import (
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFS) Open(filename string) (io.ReadWriteCloser, error) {
	fileObj, err := os.OpenFile(filename, os.O_RDWR, filesystem.RegularPerm)
	logging.ReportErr(err)

	return fileObj, err
}
