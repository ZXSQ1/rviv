package compressor

import (
	"io"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ulikunitz/xz"
)

func (compressor *XzCompressor) Compress(
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

	inObj, err := outfile.Filesys.Open(outfile.Filename, filesystem.ModeRead)

	if err != nil {
		return err
	}

	defer inObj.Close()

	xzWriter, err := xz.NewWriter(outObj)

	if err != nil {
		return err
	}

	defer xzWriter.Close()
	xzWriter.WriterConfig = compressor.Level(level)

	_, err = io.Copy(xzWriter, inObj)
	return err
}
