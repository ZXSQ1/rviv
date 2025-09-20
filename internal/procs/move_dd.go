package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/info"
)

func MoveDir(srcdir, destdir config.Path, verbose bool) error {
	if err := IsDir(srcdir); err != nil {
		return err
	}

	replaceOnCopy := false

	if err := IsExist(destdir, nil); err != nil {
		replaceOnCopy = true
	} else if err := IsDir(destdir); err != nil {
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
			Filenames: []config.Path{CopyPath(destdir, destFilename)},
			Parent:    true,
			Verbose:   verbose,
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

		destEntry := CopyPath(destdir, destEntryFilename)
		destEntryDir := CopyPath(destEntry, filepath.Dir(destEntry.Filename))

		if IsExist(destEntry, nil) != nil {
			Mkdir(config.MkdirOpts{
				Filenames: []config.Path{destEntryDir},
				Parent:    true,
				Verbose:   verbose,
			})
		}

		if IsRegular(srcEntry) == nil && (IsExist(destEntry, nil) != nil ||
			IsRegular(destEntry) == nil) {

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

		if IsDir(srcEntry) == nil && IsExist(destEntry, nil) != nil {
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
