package processes

import "github.com/ZXSQ1/rviv/logging"

// relies on the Copy function but removes the source at the end
func Move(src, dest *Path, progress chan int) error {
	err := Copy(src, dest, progress)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	return src.Filesys.Remove(src.Filename)
}
