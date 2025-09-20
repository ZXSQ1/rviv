package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/itchyny/timefmt-go"
)

func OrganizeDate(srcs []config.Path, organizedir config.Path, datefmt string,
	verbose bool) error {

	if err := IsDir(organizedir); err != nil {
		return err
	}

	for _, src := range srcs {
		if err := IsExist(src, nil); err != nil {
			return err
		}

		stat, err := Stat(src)

		if err != nil {
			return err
		}

		date := timefmt.Format(stat.ModTime(), datefmt)

		destDir := CopyPath(
			organizedir, filepath.Join(organizedir.Filename, date),
		)

		if err := IsExist(destDir, CreateDir); err != nil {
			return err
		}

		if err := IsDir(destDir); err != nil {
			return err
		}

		if stat.IsDir() {
			err := MoveDir(src, destDir, verbose)

			if err != nil {
				return err
			}

			continue
		}

		err = MoveFiles([]config.Path{src}, destDir, verbose)

		if err != nil {
			return err
		}
	}

	return nil
}
