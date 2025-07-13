package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func MoveDir(srcdir, destdir config.Path, verbose bool) error {
	if err := CheckExists(srcdir); err != nil {
		return err
	}

	if err := CheckIsDir(srcdir); err != nil {
		return err
	}

	replaceOnCopy := false

	if err := CheckExists(destdir); err != nil {
		replaceOnCopy = true
	} else if err := CheckIsDir(destdir); err != nil {
		return err
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

		if CheckIsRegular(srcEntry) == nil && (CheckExists(destEntry) != nil ||
			CheckIsRegular(destEntry) == nil) {

			err := CopyFile(srcEntry, destEntry, verbose, func(
				src, dest string) {

				info.Text(
					true, "moving source file '%s' to destination '%s'", src, dest,
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

	err = Remove(config.RemoveOpts{
		Filenames: []config.Path{srcdir},
		Recursive: true,
		Verbose:   verbose,
	})

	if err != nil {
		return err
	}

	return nil
}
