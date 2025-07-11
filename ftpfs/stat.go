package ftpfs

import (
	"io/fs"
)

func (client *FtpFs) Stat(filename string) (fs.FileInfo, error) {
	entry, err := client.conn.GetEntry(filename)

	if err != nil {
		return nil, stderr(err)
	}

	return &FileInfo{
		filename: filename,
		entry:    entry,
		conn:     client.conn,
	}, nil
}
