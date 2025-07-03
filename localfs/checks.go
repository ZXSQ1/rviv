package localfs

import (
	"os"
	"path/filepath"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFs) IsExist(filename string) bool {
	filename = filepath.Join(local.prefix, filename)
	_, err := os.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
