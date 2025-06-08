package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

// removes a file given a path
func (local *LocalFs) Remove(filename string) error {
	err := os.Remove(filename)
	logging.ReportErr(err)

	return err
}

// removes a directory given a path
func (local *LocalFs) RemoveDir(filename string) error {
	err := os.RemoveAll(filename)
	logging.ReportErr(err)

	return err
}
