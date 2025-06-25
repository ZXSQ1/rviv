package ftpfs

import (
	"io/fs"
	"path"
	"time"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
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
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return -1
	}

	return size
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
	entry, err := info.conn.GetEntry(info.filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return time.UnixMilli(0)
	}

	return entry.Time
}

func (info *FileInfo) IsRegular() bool {
	entry, err := info.conn.GetEntry(info.filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeFile
}

func (info *FileInfo) IsDir() bool {
	entry, err := info.conn.GetEntry(info.filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeFolder
}

func (info *FileInfo) IsSymlink() bool {
	entry, err := info.conn.GetEntry(info.filename)
	err = stderr(err)
	logging.ReportErr(err)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeLink
}

func (info *FileInfo) Sys() any {
	return nil
}
