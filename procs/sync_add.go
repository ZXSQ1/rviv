package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
)

func SyncAdd(src, dest config.Path, oneway, verbose bool) error {
	if err := CheckExists(src); err != nil {
		return err
	}

	if err := CheckIsDir(src); err != nil {
		return err
	}

	if err := CheckExists(dest); err != nil {
		return err
	}

	if err := CheckIsDir(dest); err != nil {
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

	if err := CopyFiles(diffs.UniqueSrcEntries, dest, verbose); err != nil {
		return err
	}

	if !oneway {
		if err := CopyFiles(diffs.UniqueDestEntries, src, verbose); err != nil {
			return err
		}
	}

	for _, commonEntry := range diffs.SrcCommonEntries {
		srcEntry := commonEntry
		destEntry := config.Path{
			Filename: filepath.Join(dest.Filename, strings.TrimLeft(
				strings.Replace(commonEntry.Filename, src.Filename, "", 1), "/",
			)),

			Devname: commonEntry.Devname,
			Devices: commonEntry.Devices,
			Active:  commonEntry.Active,
			Fsys:    commonEntry.Fsys,
		}

		srcStat, err := CheckStat(srcEntry)

		if err != nil {
			return err
		}

		destStat, err := CheckStat(destEntry)

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
