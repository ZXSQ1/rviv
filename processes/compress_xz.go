package processes

import (
	"io"

	"github.com/ZXSQ1/rviv/logging"
	"github.com/ulikunitz/xz"
)

// compresses an archive into a xz archive give 2 paths; has not been tested;
// needs implementation of compression levels
func CompressXz(archive *Path, outfile *Path) error {
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
	xzWriter, err := xz.NewWriter(outObj)
	logging.ReportErr(err)

	if err != nil {
		return err
	}

	defer xzWriter.Close()
	_, err = io.Copy(xzWriter, inObj)
	logging.ReportErr(err)

	return err
}
