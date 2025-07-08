package localfs

import (
	"os"
	"strings"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFs) Remove(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	stat, err := local.Stat(filename)

	if err != nil {
		return os.ErrNotExist
	}

	if !stat.Mode().IsRegular() {
		return filesystem.ErrFileNotRegular
	}

	return os.Remove(filename)
}

func (local *LocalFs) RemoveDir(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	stat, err := local.Stat(filename)

	if err != nil {
		return os.ErrNotExist
	}

	if !stat.IsDir() {
		return filesystem.ErrFileNotDir
	}

	return os.RemoveAll(filename)
}
