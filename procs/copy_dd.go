package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
	"github.com/pkg/errors"
)

func CopyDir(srcdir, destdir config.Path, verbose bool) error {
	srcDirStat, err := srcdir.Fsys.Stat(srcdir.Filename)

	if err != nil {
		return errors.Wrap(ErrStat, ShowPath(srcdir))
	}

	if !srcDirStat.IsDir() {
		return errors.Wrap(ErrFileNotDir, ShowPath(srcdir))
	}

	replaceOnCopy := true

	if destdir.Fsys.IsExist(destdir.Filename) {
		replaceOnCopy = false
		destDirStat, err := destdir.Fsys.Stat(destdir.Filename)

		if err != nil {
			return errors.Wrap(ErrStat, ShowPath(destdir))
		}

		if !destDirStat.IsDir() {
			return errors.Wrap(ErrFileNotDir, ShowPath(destdir))
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
		destFilename := filepath.Join(
			destdir.Filename, filepath.Base(srcdir.Filename),
		)

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
			err := Mkdir(config.MkdirOpts{
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

			if err != nil {
				return err
			}
		}

		srcEntryStat, err := srcEntry.Fsys.Stat(srcEntry.Filename)

		if err != nil {
			return errors.Wrap(ErrStat, ShowPath(srcEntry))
		}

		destEntryExists := destEntry.Fsys.IsExist(destEntry.Filename)
		destEntryStat, _ := destEntry.Fsys.Stat(destEntry.Filename)

		if srcEntryStat.IsDir() && (!destEntryExists ||
			destEntryStat.Mode().IsRegular()) {

			err := CopyFile(srcEntry, destEntry, verbose, func(
				src, dest string) {

				info.Text(
					true, "copying source file '%s' to destination '%s'", src, dest,
				)
			})

			if err != nil {
				return err
			}

			continue
		}

		if CheckIsDir(srcEntry) == nil && CheckExists(destEntry) != nil {
			err := Mkdir(config.MkdirOpts{
				Filenames: []config.Path{destEntry},
				Parent:    false,
				Verbose:   verbose,
			})

			if err != nil {
				return err
			}

			continue
		}
	}

	return nil
}
