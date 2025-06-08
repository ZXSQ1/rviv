package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFS) Stat(filename string) (os.FileInfo, error) {
	info, err := os.Stat(filename)
	logging.ReportErr(err)

	return info, err
}
