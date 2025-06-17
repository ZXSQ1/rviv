package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFs) Remove(filename string) error {
	err := os.Remove(filename)
	logging.ReportErr(err)

	return err
}

func (local *LocalFs) RemoveDir(filename string) error {
	err := os.RemoveAll(filename)
	logging.ReportErr(err)

	return err
}
