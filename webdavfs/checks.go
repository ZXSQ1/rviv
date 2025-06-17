package webdavfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (client *WebDavFs) IsExist(filename string) bool {
	_, err := client.conn.Stat(filename)
	logging.ReportErr(err)

	return !os.IsNotExist(err)
}
