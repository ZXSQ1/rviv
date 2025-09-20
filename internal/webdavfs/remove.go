package webdavfs

import "github.com/ZXSQ1/rviv/internal/filesystem"

func (client *WebDavFs) RemoveFile(filename string) error {
	if !client.IsExist(filename) {
		return filesystem.ErrNotExist
	}

	if stat, _ := client.Stat(filename); !stat.Mode().IsRegular() {
		return filesystem.ErrFileNotRegular
	}

	return client.conn.Remove(filename)
}

func (client *WebDavFs) RemoveDir(filename string) error {
	if !client.IsExist(filename) {
		return filesystem.ErrNotExist
	}

	if stat, _ := client.Stat(filename); !stat.IsDir() {
		return filesystem.ErrFileNotDir
	}

	return client.conn.RemoveAll(filename)
}
