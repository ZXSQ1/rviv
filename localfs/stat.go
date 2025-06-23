package localfs

import (
	"io/fs"
	"os"
)

func (local *LocalFs) Stat(filename string) (fs.FileInfo, error) {
	_, err := os.Stat(filename)

	if err != nil {
		return nil, err
	}

	return &FileInfo{
		filename: filename,
	}, nil
}
