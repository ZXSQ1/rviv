package compressor

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/dsnet/compress/bzip2"
)

func (compressor *Bz2Compressor) Compress(
	outfile config.Path, level CompressionLevel) error {

	if !outfile.Fsys.IsExist(outfile.Filename) {
		if err := outfile.Fsys.CreateFile(outfile.Filename); err != nil {
			return err
		}
	}

	outObj, err := outfile.Fsys.Open(outfile.Filename, filesystem.ModeWrite)

	if err != nil {
		return err
	}

	defer outObj.Close()

	inObj, err := outfile.Fsys.Open(
		compressor.filename.Filename, filesystem.ModeRead)

	if err != nil {
		return err
	}

	defer inObj.Close()

	bz2Writer, err := bzip2.NewWriter(outObj, &bzip2.WriterConfig{
		Level: int(level),
	})

	if err != nil {
		return err
	}

	defer bz2Writer.Close()

	_, err = io.Copy(bz2Writer, inObj)
	return err
}
