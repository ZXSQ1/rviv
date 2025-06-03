package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (local *LocalFS) Create(filename string) error {
	fileObj, err := os.Create(filename)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

func (local *LocalFS) CreateDir(filename string) error {
	return os.Mkdir(filename, filesystem.DirPerm)
}
