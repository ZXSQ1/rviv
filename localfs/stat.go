package localfs

import (
	"io/fs"
	"os"
	"strings"
)

func (local *LocalFs) Stat(filename string) (fs.FileInfo, error) {
	filename = strings.TrimLeft(filename, "/")
	stat, err := os.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		stat:     stat,
		filename: filename,
	}, nil
}
