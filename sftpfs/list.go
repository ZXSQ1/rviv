package sftpfs

import "github.com/ZXSQ1/rviv/logging"

// gives a listing of the entries in a directory given a path
func (client *SFtpFs) ListDir(filename string) ([]string, error) {
	resultEntries := []string{}
	entries, err := client.conn.ReadDir(filename)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		resultEntries = append(resultEntries, filename+"/"+entry.Name())
	}

	return resultEntries, nil
}
