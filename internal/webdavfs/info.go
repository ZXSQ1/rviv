package webdavfs

import (
	"io/fs"
	"path/filepath"
	"time"

	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/studio-b12/gowebdav"
)

type FileInfo struct {
	conn     *gowebdav.Client
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
	mode := filesystem.PermRegular

	if info.IsDir() {
		mode = fs.ModeDir | filesystem.PermDir
	} else if info.IsSymlink() {
		mode = fs.ModeSymlink | filesystem.PermLink
	}

	return mode
}

func (info *FileInfo) ModTime() time.Time {
	return info.stat.ModTime()
}

func (info *FileInfo) IsDir() bool {
	return info.stat.IsDir()
}

func (info *FileInfo) IsSymlink() bool {
	return !(info.stat.IsDir() || info.stat.Mode().IsRegular())
}

func (info *FileInfo) Sys() any {
	return nil
}
