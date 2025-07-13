package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func CheckExists(file config.Path) error {
	if !file.Fsys.IsExist(file.Filename) {
		return info.Error("file '%s' does not exist", ShowPath(file))
	}

	return nil
}

func CheckExistsCreate(file config.Path) error {
	if CheckExists(file) != nil {
		if file.Fsys.Create(file.Filename) != nil {
			return info.Error(
				"unable to create file '%s'", ShowPath(file),
			)
		}
	}

	return nil
}

func CheckExistsCreateDir(file config.Path) error {
	if CheckExists(file) != nil {
		if file.Fsys.CreateDir(file.Filename) != nil {
			return info.Error(
				"unable to create directory '%s'", ShowPath(file),
			)
		}
	}

	return nil
}
