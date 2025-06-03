package processes

import (
	"compress/gzip"
	"io"
)

func CompressGz(archive *Path, outfile *Path) error {
	inObj, err := archive.Filesys.Open(archive.Filename)

	if err != nil {
		return err
	}

	defer inObj.Close()
	outObj, err := archive.Filesys.Open(outfile.Filename)

	if err != nil {
		return err
	}

	defer outObj.Close()
	gzWriter := gzip.NewWriter(outObj)
	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, inObj)
	return err
}
