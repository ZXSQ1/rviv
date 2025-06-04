package ftpfs

import (
	"path"

	"github.com/jlaffaye/ftp"
)

func (client *FtpFs) IsExist(filename string) bool {
	_, err := client.conn.FileSize(filename)
	return status(err) != ftp.StatusFileActionIgnored
}

func (client *FtpFs) GetEntry(filename string) (*ftp.Entry, error) {
	entries, err := client.conn.List(path.Dir(filename))

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.Name == path.Base(filename) {
			return entry, nil
		}
	}

	return nil, nil
}
