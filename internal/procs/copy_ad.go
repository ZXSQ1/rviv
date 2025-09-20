package procs

import "github.com/ZXSQ1/rviv/internal/config"

func CopyAll(srcs []config.Path, dest config.Path, verbose bool) error {
	srcEntries, err := Glob(config.GlobOpts{
		Entries: srcs,
		Verbose: verbose,
	})

	if err != nil {
		return err
	}

	if err := IsDir(dest); err != nil {
		return err
	}

	for _, srcEntry := range srcEntries {
		var err error

		if IsDir(srcEntry) == nil {
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
