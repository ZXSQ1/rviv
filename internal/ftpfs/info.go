package ftpfs

import (
	"io/fs"
	"path"
	"time"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/jlaffaye/ftp"
)

type FileInfo struct {
	conn     *ftp.ServerConn
	entry    *ftp.Entry
	filename string
}

func (info *FileInfo) Name() string {
	return path.Base(info.filename)
}

func (info *FileInfo) Size() int64 {
	if info.entry.Type != ftp.EntryTypeFile {
		return -1
	}

	return int64(info.entry.Size)
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
	return info.entry.Time
}

func (info *FileInfo) IsRegular() bool {
	return info.entry.Type == ftp.EntryTypeFile
}

func (info *FileInfo) IsDir() bool {
	return info.entry.Type == ftp.EntryTypeFolder
}

func (info *FileInfo) IsSymlink() bool {
	return info.entry.Type == ftp.EntryTypeLink
}

func (info *FileInfo) Sys() any {
	return nil
}
