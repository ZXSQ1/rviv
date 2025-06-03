package processes

import (
	"io"

	"github.com/ulikunitz/xz"
)

func CompressXz(archive *Path, outfile *Path) error {
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
	xzWriter, err := xz.NewWriter(outObj)

	if err != nil {
		return err
	}

	defer xzWriter.Close()
	_, err = io.Copy(xzWriter, inObj)
	return err
}
