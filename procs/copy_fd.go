package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func CopyFiles(srcs []config.Path, dest config.Path, verbose bool) error {
	srcEntries, err := Glob(config.GlobOpts{
		Entries: srcs,
		Verbose: verbose,
	})

	if err != nil {
		return err
	}

	if !dest.Fsys.IsExist(dest.Filename) {
		if dest.Fsys.CreateDir(dest.Filename) != nil {
			return info.Error("unable to create directory '%s'", ShowPath(dest))
		}
	} else {
		stat, err := dest.Fsys.Stat(dest.Filename)

		if err != nil {
			return info.Error("unable to stat '%s'", ShowPath(dest))
		}

		if !stat.IsDir() {
			return info.Error(
				"destination file '%s' is not a directory", ShowPath(dest),
			)
		}
	}

	for _, srcEntry := range srcEntries {
		if !srcEntry.Fsys.IsExist(srcEntry.Filename) {
			return info.Error("source file '%s' does not exist", ShowPath(
				srcEntry))
		} else {
			stat, err := srcEntry.Fsys.Stat(srcEntry.Filename)

			if err != nil {
				return info.Error("unable to stat file '%s'", ShowPath(srcEntry))
			}

			if !stat.Mode().IsRegular() {
				return info.Error("file '%s' not a regular file", ShowPath(
					srcEntry))
			}
		}

		destEntry := config.Path{
			Filename: dest.Filename + "/" + filepath.Base(srcEntry.Filename),
			Devname:  dest.Devname,
			Devices:  dest.Devices,
			Active:   dest.Active,
			Fsys:     dest.Fsys,
		}

		if !destEntry.Fsys.IsExist(destEntry.Filename) {
			if destEntry.Fsys.Create(destEntry.Filename) != nil {
				return info.Error("unable to create regular file '%s'", ShowPath(
					destEntry))
			}
		} else {
			stat, err := destEntry.Fsys.Stat(destEntry.Filename)

			if err != nil {
				return info.Error("unable to stat file '%s'", ShowPath(destEntry))
			}

			if !stat.Mode().IsRegular() {
				return info.Error("file '%s' not a regular file", ShowPath(
					destEntry))
			}
		}

		err := CopyFile(
			srcEntry, destEntry, verbose,
			func(src, dest string) {
				info.Text(
					true, "copying source file '%s' to destination '%s'", src, dest,
				)
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
