package localfs

import (
	"io/fs"
	"strings"
)

func (local *LocalFs) Stat(filename string) (fs.FileInfo, error) {
	filename = strings.TrimLeft(filename, "/")
	stat, err := local.fsys.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		stat:     stat,
		filename: filename,
	}, nil
}
