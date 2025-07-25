package localfs

import (
	"strings"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFs) CreateFile(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	_, err := local.Stat(filename)

	if err == nil {
		return filesystem.ErrExist
	}

	fileObj, err := local.fsys.Create(filename)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

func (local *LocalFs) CreateDir(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	return local.fsys.Mkdir(filename, filesystem.PermDir)
}
