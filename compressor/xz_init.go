package compressor

import (
	"github.com/ZXSQ1/rviv/filesystem"
)

type XzCompressor struct {
	filename filesystem.Path
}

func NewXzCompressor(filename filesystem.Path) (Compressor, error) {
	if !filename.Filesys.IsExist(filename.Filename) {
		return nil, filesystem.ErrNotExist
	}

	stat, err := filename.Filesys.Stat(filename.Filename)

	if err != nil {
		return nil, err
	}

	if !stat.Mode().IsRegular() {
		return nil, filesystem.ErrFileNotRegular
	}

	return &XzCompressor{filename: filename}, nil
}
