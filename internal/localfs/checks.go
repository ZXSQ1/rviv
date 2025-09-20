package localfs

import (
	"os"
	"strings"

	"github.com/ZXSQ1/rviv/internal/logging"
)

func (local *LocalFs) IsExist(filename string) bool {
	filename = strings.TrimLeft(filename, "/")
	_, err := local.fsys.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
