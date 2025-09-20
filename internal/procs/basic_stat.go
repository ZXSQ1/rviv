package procs

import (
	"io/fs"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/pkg/errors"
)

func Stat(filename config.Path) (fs.FileInfo, error) {
	if IsExist(filename, nil) != nil {
		return nil, errors.Wrap(filesystem.ErrNotExist, ShowPath(filename))
	}

	stat, err := filename.Fsys.Stat(filename.Filename)

	if err != nil {
		return nil, errors.Wrap(filesystem.ErrStat, ShowPath(filename))
	}

	return stat, nil
}
