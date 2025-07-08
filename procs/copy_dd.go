package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func CopyDir(srcdir, destdir config.Path, verbose bool) error {
	if !srcdir.Fsys.IsExist(srcdir.Filename) {
		return info.Error(
			"source directory '%s' does not exist", ShowPath(srcdir),
		)
	} else {
		stat, err := srcdir.Fsys.Stat(srcdir.Filename)

		if err != nil {
			return info.Error(
				"unable to stat '%s'", ShowPath(srcdir),
			)
		}

		if !stat.IsDir() {
			return info.Error(
				"source file '%s' is not a directory", ShowPath(srcdir),
			)
		}
	}

	replaceOnCopy := false

	if !destdir.Fsys.IsExist(destdir.Filename) {
		replaceOnCopy = true
	} else {
		stat, err := destdir.Fsys.Stat(destdir.Filename)

		if err != nil {
			return info.Error(
				"unable to stat '%s'", ShowPath(destdir),
			)
		}

		if !stat.IsDir() {
			return info.Error(
				"destination file '%s' is not a directory", ShowPath(destdir),
			)
		}
	}

	srcEntries, err := ListDir(config.ListOpts{
		Filename:  srcdir,
		Recursive: true,
		Verbose:   verbose,
	})

	if err != nil {
		return err
	}

	if len(srcEntries) == 0 {
		destFilename := filepath.Join(destdir.Filename, filepath.Base(
			srcdir.Filename))

		if replaceOnCopy {
			destFilename = destdir.Filename
		}

		err := Mkdir(config.MkdirOpts{
			Filenames: []config.Path{
				{
					Filename: destFilename,
					Devname:  destdir.Devname,
					Devices:  destdir.Devices,
					Active:   destdir.Active,
					Fsys:     destdir.Fsys,
				},
			},

			Parent:  true,
			Verbose: verbose,
		})

		if err != nil {
			return err
		}
	}

	for _, srcEntry := range srcEntries {
		destEntryFilename := filepath.Join(
			destdir.Filename, filepath.Base(srcdir.Filename), strings.Replace(
				srcEntry.Filename, srcdir.Filename, "", 1,
			),
		)

		if replaceOnCopy {
			destEntryFilename = filepath.Join(
				destdir.Filename, strings.Replace(
					srcEntry.Filename, srcdir.Filename, "", 1,
				),
			)
		}

		destEntry := config.Path{
			Filename: destEntryFilename,
			Devname:  destdir.Devname,
			Devices:  destdir.Devices,
			Active:   destdir.Active,
			Fsys:     destdir.Fsys,
		}

		if !destEntry.Fsys.IsExist(filepath.Dir(destEntry.Filename)) {
			Mkdir(config.MkdirOpts{
				Filenames: []config.Path{
					{
						Filename: filepath.Dir(destEntry.Filename),
						Devices:  destEntry.Devices,
						Devname:  destEntry.Devname,
						Active:   destEntry.Active,
						Fsys:     destEntry.Fsys,
					},
				},

				Parent:  true,
				Verbose: verbose,
			})
		}

		srcStat, err := srcEntry.Fsys.Stat(srcEntry.Filename)

		if err != nil {
			return info.Error("unable to stat '%s'", ShowPath(srcEntry))
		}

		if !destEntry.Fsys.IsExist(destEntry.Filename) {
			if srcStat.IsDir() {
				Mkdir(config.MkdirOpts{
					Filenames: []config.Path{destEntry},
					Parent:    false,
					Verbose:   verbose,
				})
			} else {
				err := CopyFile(srcEntry, destEntry, verbose, func(
					src, dest string) {

					info.Text(
						true, "copying source file '%s' to destination '%s'",
						src, dest,
					)
				})

				if err != nil {
					return err
				}
			}
		} else {
			destStat, err := destEntry.Fsys.Stat(destEntry.Filename)

			if err != nil {
				info.Error("unable to stat '%s'", ShowPath(destEntry))
			}

			if srcStat.IsDir() != destStat.IsDir() {
				return info.Error(
					"destination file '%s' is not of the same file "+
						"type as the source file '%s'", ShowPath(srcEntry),
					ShowPath(destEntry),
				)
			} else if srcStat.Mode().IsRegular() {
				err := CopyFile(srcEntry, destEntry, verbose, func(
					src, dest string) {

					info.Text(
						true, "copying source file '%s' to destination '%s'",
						src, dest,
					)
				})

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
