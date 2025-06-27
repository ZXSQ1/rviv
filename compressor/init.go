package compressor

import "github.com/ZXSQ1/rviv/filesystem"

type Compressor struct {
	filename filesystem.Path
}

func NewCompressor(filename filesystem.Path) (Compressor, error) {
	return Compressor{}, nil
}
