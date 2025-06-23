package localfs

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/ZXSQ1/rviv/logging"
)

type FileInfo struct {
	filename string
}

func (info *FileInfo) Name() string {
	return filepath.Base(info.filename)
}

func (info *FileInfo) Size() int64 {
	stat, err := os.Stat(info.filename)
	logging.ReportErr(err)

	if stat.IsDir() {
		return -1
	}

	return stat.Size()
}

func (info *FileInfo) Mode() fs.FileMode {
	stat, err := os.Stat(info.filename)
	logging.ReportErr(err)

	return stat.Mode()
}

func (info *FileInfo) ModTime() time.Time {
	stat, err := os.Stat(info.filename)
	logging.ReportErr(err)

	return stat.ModTime()
}

func (info *FileInfo) IsDir() bool {
	stat, err := os.Stat(info.filename)
	logging.ReportErr(err)

	return stat.IsDir()
}

func (info *FileInfo) Sys() any {
	return nil
}
