package procs

import (
	"io"
	"io/fs"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/ZXSQ1/rviv/info"
)

func CheckStat(file config.Path) (fs.FileInfo, error) {
	stat, err := file.Fsys.Stat(file.Filename)

	if err != nil {
		return nil, info.Error("unable to stat '%s'", ShowPath(file))
	}

	return stat, nil
}

func CheckOpen(file config.Path, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	modeString := ""
	obj, err := file.Fsys.Open(file.Filename, mode)

	if mode == filesystem.ModeWrite {
		modeString = "reading"
	} else if mode == filesystem.ModeRead {
		modeString = "writing"
	}

	if err != nil {
		return nil, info.Error(
			"unable to open '%s' for %s", ShowPath(file), modeString,
		)
	}

	return obj, nil
}
