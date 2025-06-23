package sftpfs

import (
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *SFtpFs) Remove(filename string) error {
	if !client.IsExist(filename) {
		return os.ErrNotExist
	}

	stat, _ := client.Stat(filename)

	if !stat.Mode().IsRegular() {
		return filesystem.ErrFileNotRegular
	}

	return client.conn.Remove(filename)
}

func (client *SFtpFs) RemoveDir(filename string) error {
	if !client.IsExist(filename) {
		return os.ErrNotExist
	}

	stat, _ := client.Stat(filename)

	if !stat.IsDir() {
		return filesystem.ErrFileNotDir
	}

	return client.conn.RemoveAll(filename)
}
