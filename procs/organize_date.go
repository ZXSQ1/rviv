package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/config"
	"github.com/itchyny/timefmt-go"
)

func OrganizeDate(srcs []config.Path, organizedir config.Path, datefmt string,
	verbose bool) error {

	if err := CheckExists(organizedir); err != nil {
		return err
	}

	if err := CheckIsDir(organizedir); err != nil {
		return err
	}

	for _, src := range srcs {
		if err := CheckExists(src); err != nil {
			return err
		}

		stat, err := CheckStat(src)

		if err != nil {
			return err
		}

		date := timefmt.Format(stat.ModTime(), datefmt)

		destDir := config.Path{
			Filename: filepath.Join(organizedir.Filename, date),
			Devname:  organizedir.Devname,
			Devices:  organizedir.Devices,
			Active:   organizedir.Active,
			Fsys:     organizedir.Fsys,
		}

		if err := CheckExistsCreateDir(destDir); err != nil {
			return err
		}

		if err := CheckIsDir(destDir); err != nil {
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
