package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
)

// creates a file given a path
func (local *LocalFs) Create(filename string) error {
	fileObj, err := os.Create(filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	return fileObj.Close()
}

// creates a directory given a path
func (local *LocalFs) CreateDir(filename string) error {
	err := os.Mkdir(filename, filesystem.DirPerm)
	logging.ReportErr(err)

	return err
}
