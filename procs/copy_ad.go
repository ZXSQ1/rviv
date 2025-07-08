package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func CopyAll(srcs []config.Path, dest config.Path, verbose bool) error {
	srcEntries, err := Glob(config.GlobOpts{
		Entries: srcs,
		Verbose: verbose,
	})

	if err != nil {
		return err
	}

	if !dest.Fsys.IsExist(dest.Filename) {
		return info.Error(
			"destination directory '%s' does not exist", ShowPath(dest),
		)
	} else {
		stat, err := dest.Fsys.Stat(dest.Filename)

		if err != nil {
			return info.Error("unable to stat '%s'", ShowPath(dest))
		}

		if !stat.IsDir() {
			return info.Error(
				"file '%s' is not a directory", ShowPath(dest),
			)
		}
	}

	for _, srcEntry := range srcEntries {
		srcStat, err := srcEntry.Fsys.Stat(srcEntry.Filename)

		if err != nil {
			return info.Error("unable to stat '%s'", ShowPath(dest))
		}

		if srcStat.IsDir() {
			if err = CopyDir(srcEntry, dest, verbose); err != nil {
				return err
			}
		} else {
			err = CopyFiles([]config.Path{srcEntry}, dest, verbose)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
