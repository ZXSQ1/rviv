package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/pkg/errors"
)

func CopyAll(srcs []config.Path, dest config.Path, verbose bool) error {
	srcEntries, err := Glob(config.GlobOpts{
		Entries: srcs,
		Verbose: verbose,
	})

	if err != nil {
		return err
	}

	destStat, err := dest.Fsys.Stat(dest.Filename)

	if err != nil {
		return errors.Wrap(ErrStat, ShowPath(dest))
	}

	if !destStat.IsDir() {
		return errors.Wrap(ErrFileNotDir, ShowPath(dest))
	}

	for _, srcEntry := range srcEntries {
		var err error

		srcEntryStat, err := srcEntry.Fsys.Stat(srcEntry.Filename)

		if err != nil {
			return errors.Wrap(ErrStat, ShowPath(srcEntry))
		}

		if srcEntryStat.IsDir() {
			err = CopyDir(srcEntry, dest, verbose)
		} else {
			err = CopyFiles([]config.Path{srcEntry}, dest, verbose)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
