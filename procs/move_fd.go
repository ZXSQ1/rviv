package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func MoveFiles(srcs []config.Path, dest config.Path, verbose bool) error {
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
		if err := AssertExists(srcEntry); err != nil {
			return err
		}

		if err := AssertIsRegular(srcEntry); err != nil {
			return err
		}

		destEntry := config.Path{
			Filename: dest.Filename + "/" + filepath.Base(srcEntry.Filename),
			Devname:  dest.Devname,
			Devices:  dest.Devices,
			Active:   dest.Active,
			Fsys:     dest.Fsys,
		}

		if err := AssertExistsCreate(destEntry); err != nil {
			return err
		}

		if err := AssertIsRegular(destEntry); err != nil {
			return err
		}

		err := MoveFile(
			srcEntry, destEntry, verbose, func(src, dest string) {
				info.Text(
					true, "moving source file '%s' to destination '%s'", src, dest,
				)
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
