package ftpfs

import "github.com/ZXSQ1/rviv/logging"

// gives a listing of the entries in a directory given its path
func (client *FtpFs) ListDir(filename string) ([]string, error) {
	entries, err := client.conn.NameList(filename)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	return entries, nil
}
