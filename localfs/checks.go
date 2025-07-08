package localfs

import (
	"os"
	"strings"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFs) IsExist(filename string) bool {
	filename = strings.TrimLeft(filename, "/")
	_, err := os.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
