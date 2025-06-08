package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFS) IsExist(filename string) bool {
	_, err := os.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
