package processes

import (
	"compress/gzip"
	"io"

	"github.com/ZXSQ1/rviv/logging"
)

func CompressGz(archive *Path, outfile *Path) error {
	inObj, err := archive.Filesys.Open(archive.Filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer inObj.Close()
	outObj, err := archive.Filesys.Open(outfile.Filename)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer outObj.Close()
	gzWriter := gzip.NewWriter(outObj)
	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, inObj)
	logging.ReportErr(err)

	return err
}
