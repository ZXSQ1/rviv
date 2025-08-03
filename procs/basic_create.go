package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/pkg/errors"
)

func CreateFile(filename config.Path) error {
	if IsExist(filename) == nil {
		return errors.Wrap(filesystem.ErrExist, ShowPath(filename))
	}

	parent := config.Path{
		Filename: filepath.Dir(filename.Filename),
		Devices:  filename.Devices,
		Devname:  filename.Devname,
		Active:   filename.Active,
		Fsys:     filename.Fsys,
	}

	if IsExist(parent) != nil {
		if err := CreateDir(parent); err != nil {
			return err
		}
	}

	if err := IsDir(parent); err != nil {
		return err
	}

	if filename.Fsys.CreateFile(filename.Filename) != nil {
		return errors.Wrap(filesystem.ErrCreateFile, ShowPath(filename))
	}

	return nil
}

func CreateDir(filename config.Path) error {
	if IsExist(filename) == nil {
		return errors.Wrap(filesystem.ErrExist, ShowPath(filename))
	}

	parent := config.Path{
		Filename: filepath.Dir(filename.Filename),
		Devices:  filename.Devices,
		Devname:  filename.Devname,
		Active:   filename.Active,
		Fsys:     filename.Fsys,
	}

	if IsExist(parent) != nil {
		if err := CreateDir(parent); err != nil {
			return err
		}
	}

	if err := IsDir(parent); err != nil {
		return err
	}

	if filename.Fsys.CreateDir(filename.Filename) != nil {
		return errors.Wrap(filesystem.ErrCreateDir, ShowPath(filename))
	}

	return nil
}
