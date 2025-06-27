package compressor

import "github.com/ZXSQ1/rviv/filesystem"

type Bz2Compressor struct {
	filename filesystem.Path
}

func NewBz2Compressor(filename filesystem.Path) (Compressor, error) {
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

	return &Bz2Compressor{filename: filename}, nil
}
