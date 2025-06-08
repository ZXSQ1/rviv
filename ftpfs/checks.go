package ftpfs

import (
	"strings"

	"github.com/ZXSQ1/rviv/logging"
)

// checks if the file with the given path exists; uses the returned error from
// querying the file size of the file; checks for "no such file or directory" at
// the end of the error message
func (client *FtpFs) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	logging.ReportErr(err)

	if err != nil {
		return !strings.HasSuffix(err.Error(), "no such file or directory")
	}

	return true
}
