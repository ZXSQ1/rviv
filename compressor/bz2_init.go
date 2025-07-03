package compressor

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
)

type Bz2Compressor struct {
	filename config.Path
}

func NewBz2Compressor(filename config.Path) (Compressor, error) {
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

	return &Bz2Compressor{filename: filename}, nil
}
