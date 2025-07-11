package webdavfs

import (
	"io/fs"
)

func (client *WebDavFs) Stat(filename string) (fs.FileInfo, error) {
	stat, err := client.conn.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		filename: filename,
		stat:     stat,
		conn:     client.conn,
	}, nil
}
