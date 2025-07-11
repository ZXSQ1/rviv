package procs

import (
	"github.com/ZXSQ1/rviv/config"
)

func MoveAll(srcs []config.Path, dest config.Path, verbose bool) error {
	srcEntries, err := Glob(config.GlobOpts{
		Entries: srcs,
		Verbose: verbose,
	})

	if err != nil {
		return err
	}

	if err := AssertExists(dest); err != nil {
		return err
	}

	if err := AssertIsDir(dest); err != nil {
		return err
	}

	for _, srcEntry := range srcEntries {
		if AssertIsDir(srcEntry) == nil {
			if err = MoveDir(srcEntry, dest, verbose); err != nil {
				return err
			}

			continue
		}

		err = MoveFiles([]config.Path{srcEntry}, dest, verbose)

		if err != nil {
			return err
		}
	}

	return nil
}
