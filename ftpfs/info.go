package ftpfs

import (
	"io/fs"
	"path"
	"time"

	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/logging"
	"github.com/jlaffaye/ftp"
)

// an implementation of the fs.FileInfo object that represents the os.FileInfo
// structure for FTP; returned in the Stat method
type FileInfo struct {
	conn     *ftp.ServerConn
	filename string
}

// gets the base name of the file
func (info *FileInfo) Name() string {
	return path.Base(info.filename)
}

// gets the size of the file; the size is -1 if there is an error
func (info *FileInfo) Size() int64 {
	size, err := info.conn.FileSize(info.filename)
	logging.ReportErr(err)

	if err != nil {
		return -1
	}

	return size
}

// gets the mode of the file
func (info *FileInfo) Mode() fs.FileMode {
	mode := filesystem.RegularPerm

	if info.IsDir() {
		mode = fs.ModeDir | filesystem.DirPerm
	} else if info.IsSymlink() {
		mode = fs.ModeSymlink | filesystem.LinkPerm
	}

	return mode
}

// gets the modification time of the file
func (info *FileInfo) ModTime() time.Time {
	entry, err := info.conn.GetEntry(info.filename)
	logging.ReportErr(err)

	if err != nil {
		return time.UnixMilli(0)
	}

	return entry.Time
}

// checks if the file is regular; not necessary for fs.FileInfo implementation
func (info *FileInfo) IsRegular() bool {
	entry, err := info.conn.GetEntry(info.filename)
	logging.ReportErr(err)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeFile
}

// checks if the file is a directory
func (info *FileInfo) IsDir() bool {
	entry, err := info.conn.GetEntry(info.filename)
	logging.ReportErr(err)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeFolder
}

// checks if the file is a symbolic link; not necessary for the fs.FileInfo
// implementation
func (info *FileInfo) IsSymlink() bool {
	entry, err := info.conn.GetEntry(info.filename)
	logging.ReportErr(err)

	if err != nil {
		return false
	}

	return entry.Type == ftp.EntryTypeLink
}

// some system settings; usually does not affect anything, and most systems just
// return nil
func (info *FileInfo) Sys() any {
	return nil
}
