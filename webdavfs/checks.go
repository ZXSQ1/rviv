package webdavfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
)

func (client *WebDavFs) IsExist(filename string) bool {
	_, err := client.conn.Stat(filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return err.Error() != filesystem.ErrNotExist.Error()
	}

	return true
}
