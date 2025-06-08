package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFS) Create(filename string) error {
	fileObj, err := os.Create(filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

func (local *LocalFS) CreateDir(filename string) error {
	err := os.Mkdir(filename, filesystem.DirPerm)
	logging.ReportErr(err)

	return err
}
