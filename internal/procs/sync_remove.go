package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/internal/config"
)

func SyncRemove(src, dest config.Path, oneway, verbose bool) error {
	if err := IsDir(src); err != nil {
		return err
	}

	if err := IsDir(dest); err != nil {
		return err
	}

	diffs, err := Differ(config.DifferOpts{
		Src:     src,
		Dest:    dest,
		Verbose: verbose,
	})

	if err != nil {
		return err
	}

	err = Remove(config.RemoveOpts{
		Filenames: diffs.UniqueDestEntries,
		Recursive: true,
		Verbose:   verbose,
	})

	if err != nil {
		return err
	}

	if !oneway {
		err = Remove(config.RemoveOpts{
			Filenames: diffs.UniqueSrcEntries,
			Recursive: true,
			Verbose:   verbose,
		})

		if err != nil {
			return err
		}
	}

	for _, commonEntry := range diffs.SrcCommonEntries {
		srcEntry := commonEntry
		destEntry := CopyPath(
			commonEntry, filepath.Join(dest.Filename, strings.TrimLeft(
				strings.Replace(commonEntry.Filename, src.Filename, "", 1), "/",
			)),
		)

		srcStat, err := Stat(srcEntry)

		if err != nil {
			return err
		}

		destStat, err := Stat(destEntry)

		if err != nil {
			return err
		}

		if !(srcStat.Size() > destStat.Size()) {
			continue
		}

		err = CopyAll([]config.Path{srcEntry}, dest, verbose)

		if err != nil {
			return err
		}
	}

	return nil
}
