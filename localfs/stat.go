package localfs

import (
	"io/fs"
	"os"
	"path/filepath"
)

func (local *LocalFs) Stat(filename string) (fs.FileInfo, error) {
	filename = filepath.Join(local.prefix, filename)
	_, err := os.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		filename: filename,
	}, nil
}
