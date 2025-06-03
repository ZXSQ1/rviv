package processes

import (
	"io"

	"github.com/dsnet/compress/bzip2"
)

func CompressBz2(archive *Path, outfile *Path) error {
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
	bz2Writer, err := bzip2.NewWriter(
		outObj, &bzip2.WriterConfig{Level: bzip2.BestCompression},
	)

	if err != nil {
		return err
	}

	defer bz2Writer.Close()

	_, err = io.Copy(bz2Writer, inObj)
	return err
}
