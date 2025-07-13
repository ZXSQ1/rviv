package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func CheckIsRegular(file config.Path) error {
	stat, err := CheckStat(file)

	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return info.Error("file '%s' not a regular file", ShowPath(file))
	}

	return nil
}

func CheckIsDir(file config.Path) error {
	stat, err := CheckStat(file)

	if err != nil {
		return err
	}

	if !stat.IsDir() {
		return info.Error("file '%s' not a directory", ShowPath(file))
	}

	return nil
}
