package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
)

func OrganizeAlpha(srcs []config.Path, organizedir config.Path,
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

		firstLetter := strings.ToUpper(string(
			filepath.Base(src.Filename)[0],
		))

		destDir := config.Path{
			Filename: filepath.Join(organizedir.Filename, firstLetter),
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

		if CheckIsDir(src) == nil {
			err := MoveDir(src, destDir, verbose)

			if err != nil {
				return err
			}

			continue
		}

		err := MoveFiles([]config.Path{src}, destDir, verbose)

		if err != nil {
			return err
		}
	}

	return nil
}
