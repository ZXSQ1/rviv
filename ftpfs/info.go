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
	filename string
}

func (info *FileInfo) Name() string {
	return path.Base(info.filename)
}

func (info *FileInfo) Size() int64 {
	size, err := info.conn.FileSize(info.filename)

	if err != nil {
		return -1
	}

	return size
}

func (info *FileInfo) Mode() fs.FileMode {
	if info == nil {
		return 0
	}

	mode := filesystem.RegularPerm

	if info.IsDir() {
		mode = fs.ModeDir | filesystem.DirPerm
	} else if info.IsSymlink() {
		mode = fs.ModeSymlink | filesystem.LinkPerm
	}

	return mode
}

func (info *FileInfo) ModTime() time.Time {
	entry, err := info.conn.GetEntry(info.filename)

	if err != nil {
		return time.UnixMilli(0)
	}

	return entry.Time
}

func (info *FileInfo) IsRegular() bool {
	entry, err := info.conn.GetEntry(info.filename)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeFile
}

func (info *FileInfo) IsDir() bool {
	entry, err := info.conn.GetEntry(info.filename)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeFolder
}

func (info *FileInfo) IsSymlink() bool {
	entry, err := info.conn.GetEntry(info.filename)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeLink
}

func (info *FileInfo) Sys() any {
	return nil
}
