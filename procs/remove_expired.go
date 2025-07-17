package procs

import (
	"path/filepath"
	"time"

	"github.com/ZXSQ1/rviv/config"
)

func RemoveExpired(parent config.Path, baseFilenameFmt string,
	expiryTime time.Time, verbose bool) error {

	entries, err := ListDir(config.ListOpts{
		Filename:  parent,
		Recursive: false,
		Verbose:   verbose,
	})

	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.Filename != filepath.Join(
			parent.Filename, config.StdPath(baseFilenameFmt),
		) {
			continue
		}

		stat, err := CheckStat(entry)

		if err != nil {
			return err
		}

		if expiryTime.Compare(stat.ModTime()) == -1 {
			continue
		}

		err = Remove(config.RemoveOpts{
			Filenames: []config.Path{entry},
			Recursive: true,
			Verbose:   verbose,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
