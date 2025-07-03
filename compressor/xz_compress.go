package compressor

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ulikunitz/xz"
)

func (compressor *XzCompressor) Compress(
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

	inObj, err := compressor.filename.Fsys.Open(
		compressor.filename.Filename, filesystem.ModeRead)

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
