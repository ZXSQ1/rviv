package localfs

import (
	"io/fs"
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFs) Stat(filename string) (fs.FileInfo, error) {
	info, err := os.Stat(filename)
	logging.ReportErr(err)

	return info, err
}
