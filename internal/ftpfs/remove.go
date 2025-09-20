package ftpfs

import (
	"github.com/ZXSQ1/rviv/filesystem"
)

func (client *FtpFs) RemoveFile(filename string) error {
	stat, err := client.Stat(filename)

	if err != nil {
		return stderr(err)
	}

	if !stat.Mode().IsRegular() {
		return filesystem.ErrFileNotRegular
	}

	return stderr(client.conn.Delete(filename))
}

func (client *FtpFs) RemoveDir(filename string) error {
	stat, err := client.Stat(filename)

	if err != nil {
		return stderr(err)
	}

	if !stat.IsDir() {
		return filesystem.ErrFileNotDir
	}

	return stderr(client.conn.RemoveDirRecur(filename))
}
