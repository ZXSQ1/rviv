package webdavfs

import (
	"io/fs"
)

func (client *WebDavFs) Stat(filename string) (fs.FileInfo, error) {
	return client.conn.Stat(filename)
}
