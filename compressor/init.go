package compressor

import (
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
)

type Compressor struct {
	filename filesystem.Path
}

func NewCompressor(filename filesystem.Path) Compressor {
	if !filename.Filesys.IsExist(filename.Filename) {
		info.Error(
			"could not use file \"%s\" as compressor source",
			filename.Filename,
		)
	}

	stat, err := filename.Filesys.Stat(filename.Filename)

	if err != nil {

	}
}
