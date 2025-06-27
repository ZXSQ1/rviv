package compressor

import (
	"compress/gzip"
	"io"

	"github.com/ZXSQ1/rviv/filesystem"
)

func (compressor *GzCompressor) Compress(
	outfile filesystem.Path, level CompressionLevel) error {

	if !outfile.Filesys.IsExist(outfile.Filename) {
		if err := outfile.Filesys.Create(outfile.Filename); err != nil {
			return err
		}
	}

	outObj, err := outfile.Filesys.Open(outfile.Filename, filesystem.ModeWrite)

	if err != nil {
		return err
	}

	defer outObj.Close()

	inObj, err := outfile.Filesys.Open(
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
