package localfs

import (
	"strings"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFs) Remove(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	stat, err := local.Stat(filename)

	if err != nil {
		return filesystem.ErrNotExist
	}

	if !stat.Mode().IsRegular() {
		return filesystem.ErrFileNotRegular
	}

	return local.fsys.Remove(filename)
}

func (local *LocalFs) RemoveDir(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	stat, err := local.Stat(filename)

	if err != nil {
		return filesystem.ErrNotExist
	}

	if !stat.IsDir() {
		return filesystem.ErrFileNotDir
	}

	return local.fsys.RemoveAll(filename)
}
