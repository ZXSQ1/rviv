package compressor

import (
	"compress/gzip"
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
)

func (compressor *GzCompressor) Compress(
	outfile config.Path, level CompressionLevel) error {

	if !outfile.Fsys.IsExist(outfile.Filename) {
		if err := outfile.Fsys.Create(outfile.Filename); err != nil {
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

	gzWriter, err := gzip.NewWriterLevel(outObj, int(level))

	if err != nil {
		return err
	}

	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, inObj)
	return err
}
