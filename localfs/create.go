package localfs

import (
	"os"
	"strings"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFs) Create(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	_, err := local.Stat(filename)

	if err == nil {
		return filesystem.ErrExist
	}

	fileObj, err := os.Create(filename)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

func (local *LocalFs) CreateDir(filename string) error {
	filename = strings.TrimLeft(filename, "/")
	return os.Mkdir(filename, filesystem.PermDir)
}
