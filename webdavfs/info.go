package webdavfs

import (
	"io/fs"
	"path/filepath"
	"time"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
	"github.com/studio-b12/gowebdav"
)

type FileInfo struct {
	conn     *gowebdav.Client
	filename string
}

func (info *FileInfo) Name() string {
	return filepath.Base(info.filename)
}

func (info *FileInfo) Size() int64 {
	stat, err := info.conn.Stat(info.filename)
	logging.ReportErr(err)

	if stat.IsDir() {
		return -1
	}

	return stat.Size()
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
	stat, err := info.conn.Stat(info.filename)
	logging.ReportErr(err)

	return stat.ModTime()
}

func (info *FileInfo) IsDir() bool {
	stat, err := info.conn.Stat(info.filename)
	logging.ReportErr(err)

	return stat.IsDir()
}

func (info *FileInfo) IsSymlink() bool {
	stat, err := info.conn.Stat(info.filename)
	logging.ReportErr(err)

	return !(stat.IsDir() || stat.Mode().IsRegular())
}

func (info *FileInfo) Sys() any {
	return nil
}
