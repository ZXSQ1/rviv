package localfs

import (
	"io/fs"
	"os"
)

func (local *LocalFs) Stat(filename string) (fs.FileInfo, error) {
	return os.Stat(filename)
}
