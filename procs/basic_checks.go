package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/pkg/errors"
)

func IsExist(filename config.Path) error {
	if !filename.Fsys.IsExist(filename.Filename) {
		return errors.Wrap(filesystem.ErrNotExist, ShowPath(filename))
	}

	return nil
}

func IsRegular(filename config.Path) error {
	if IsExist(filename) != nil {
		return errors.Wrap(filesystem.ErrNotExist, ShowPath(filename))
	}

	stat, err := Stat(filename)

	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return errors.Wrap(filesystem.ErrFileNotRegular, ShowPath(filename))
	}

	return nil
}

func IsDir(filename config.Path) error {
	if IsExist(filename) != nil {
		return errors.Wrap(filesystem.ErrNotExist, ShowPath(filename))
	}

	stat, err := Stat(filename)

	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return errors.Wrap(filesystem.ErrFileNotDir, ShowPath(filename))
	}

	return nil
}
