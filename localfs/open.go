package localfs

import (
	"io"
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
)

// opens a file given a path
func (local *LocalFs) Open(filename string) (io.ReadWriteCloser, error) {
	fileObj, err := os.OpenFile(filename, os.O_RDWR, filesystem.RegularPerm)
	logging.ReportErr(err)

	return fileObj, err
}
