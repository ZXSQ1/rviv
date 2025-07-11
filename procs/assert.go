package procs

import (
	"io/fs"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func AssertExists(file config.Path) error {
	if !file.Fsys.IsExist(file.Filename) {
		return info.Error("file '%s' does not exist", ShowPath(file))
	}

	return nil
}

func AssertExistsCreate(file config.Path) error {
	if AssertExists(file) != nil {
		if file.Fsys.Create(file.Filename) != nil {
			return info.Error(
				"unable to create file '%s'", ShowPath(file),
			)
		}
	}

	return nil
}

func AssertExistsCreateDir(file config.Path) error {
	if AssertExists(file) != nil {
		if file.Fsys.CreateDir(file.Filename) != nil {
			return info.Error(
				"unable to create directory '%s'", ShowPath(file),
			)
		}
	}

	return nil
}

func AssertStatWorks(file config.Path) (fs.FileInfo, error) {
	stat, err := file.Fsys.Stat(file.Filename)

	if err != nil {
		return nil, info.Error("unable to stat '%s'", ShowPath(file))
	}

	return stat, nil
}

func AssertIsRegular(file config.Path) error {
	stat, err := AssertStatWorks(file)

	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return info.Error("file '%s' not a regular file", ShowPath(file))
	}

	return nil
}

func AssertIsDir(file config.Path) error {
	stat, err := AssertStatWorks(file)

	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return info.Error("file '%s' not a directory", ShowPath(file))
	}

	return nil
}
