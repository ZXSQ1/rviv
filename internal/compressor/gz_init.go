package compressor

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
)

type GzCompressor struct {
	filename config.Path
}

func NewGzCompressor(filename config.Path) (Compressor, error) {
	if !filename.Fsys.IsExist(filename.Filename) {
		return nil, filesystem.ErrNotExist
	}

	stat, err := filename.Fsys.Stat(filename.Filename)

	if err != nil {
		return nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, filesystem.ErrFileNotRegular
	}

	return &GzCompressor{filename: filename}, nil
}
