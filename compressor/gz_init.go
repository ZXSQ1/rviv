package compressor

import "github.com/ZXSQ1/rviv/filesystem"

type GzCompressor struct {
	filename filesystem.Path
}

func NewGzCompressor(filename filesystem.Path) (Compressor, error) {
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

	return &GzCompressor{filename: filename}, nil
}
