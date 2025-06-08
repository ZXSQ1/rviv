package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

// checks if the file exists given a path
func (local *LocalFs) IsExist(filename string) bool {
	_, err := os.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
