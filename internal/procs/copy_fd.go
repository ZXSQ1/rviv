package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/info"
)

func CopyFiles(srcs []config.Path, dest config.Path, verbose bool) error {
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
		if err := IsRegular(srcEntry); err != nil {
			return err
		}

		destEntry := CopyPath(
			dest, dest.Filename+"/"+filepath.Base(srcEntry.Filename),
		)

		if err := IsExist(destEntry, CreateFile); err != nil {
			return err
		}

		if err := IsRegular(destEntry); err != nil {
			return err
		}

		err := CopyFile(
			srcEntry, destEntry, verbose, func(src, dest string) {
				info.Text(
					true, "'%s' => '%s'", src, dest,
				)
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
