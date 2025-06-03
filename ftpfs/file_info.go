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

func (info FileInfo) Name() string {
	return path.Base(info.filename)
}

func (info FileInfo) Size() int64 {
	size, err := info.conn.FileSize(info.filename)

	if err != nil {
		return -1
	}

	return size
}

func (info FileInfo) Mode() fs.FileMode {
	var mode fs.FileMode

	if info.IsDir() {
		mode |= fs.ModeDir
	} else if !info.IsDir() && !info.IsRegular() {
		mode |= fs.ModeSymlink
	}

	if info.IsDir() {
		mode |= filesystem.DirPerm
	} else if !info.IsRegular() && !info.IsDir() {
		mode |= filesystem.LinkPerm
	} else {
		mode |= filesystem.RegularPerm
	}

	return mode
}

func (info FileInfo) ModTime() time.Time {
	modTime, _ := info.conn.GetTime(info.filename)
	return modTime
}

func (info FileInfo) IsRegular() bool {
	entry, _ := info.conn.GetEntry(info.filename)
	return entry.Type == ftp.EntryTypeFile
}

func (info FileInfo) IsDir() bool {
	entry, _ := info.conn.GetEntry(info.filename)
	return entry.Type == ftp.EntryTypeFolder
}

func (info FileInfo) Sys() any {
	return nil
}
