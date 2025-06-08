package processes

import (
	"errors"
	"io"

	"github.com/ZXSQ1/rviv/logging"
)

func Copy(src, dest *Path, progress chan int) error {
	srcfs, destfs := src.Filesys, dest.Filesys
	srcobj, err := srcfs.Open(src.Filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer srcobj.Close()
	destobj, err := destfs.Open(dest.Filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer destobj.Close()

	for {
		buffer := make([]byte, BufferSize)
		n, err := srcobj.Read(buffer)
		logging.ReportErr(err)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		_, err = destobj.Write(buffer[:n])
		logging.ReportErr(err)

		if err != nil {
			return err
		}

		if progress != nil {
			progress <- n
		}
	}

	return nil
}
