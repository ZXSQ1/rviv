package localfs

import (
	"io/fs"
	"path/filepath"
	"time"
)

type FileInfo struct {
	stat     fs.FileInfo
	filename string
}

func (info *FileInfo) Name() string {
	return filepath.Base(info.filename)
}

func (info *FileInfo) Size() int64 {
	if info.stat.IsDir() {
		return -1
	}

	return info.stat.Size()
}

func (info *FileInfo) Mode() fs.FileMode {
	return info.stat.Mode()
}

func (info *FileInfo) ModTime() time.Time {
	return info.stat.ModTime()
}

func (info *FileInfo) IsDir() bool {
	return info.stat.IsDir()
}

func (info *FileInfo) Sys() any {
	return nil
}
