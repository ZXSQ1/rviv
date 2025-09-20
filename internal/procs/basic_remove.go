package procs

import (
	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/filesystem"
	"github.com/pkg/errors"
)

func RemoveFile(filename config.Path) error {
	if IsRegular(filename) != nil {
		return errors.Wrap(filesystem.ErrFileNotRegular, ShowPath(filename))
	}

	if filename.Fsys.RemoveFile(filename.Filename) != nil {
		return errors.Wrap(filesystem.ErrRemoveFile, ShowPath(filename))
	}

	return nil
}

func RemoveDir(filename config.Path) error {
	if IsDir(filename) != nil {
		return errors.Wrap(filesystem.ErrFileNotDir, ShowPath(filename))
	}

	if filename.Fsys.RemoveDir(filename.Filename) != nil {
		return errors.Wrap(filesystem.ErrRemoveFile, ShowPath(filename))
	}

	return nil
}
