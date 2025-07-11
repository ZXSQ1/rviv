package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
)

func OrganizeExt(srcs []config.Path, organizedir config.Path,
	verbose bool) error {

	if err := AssertExists(organizedir); err != nil {
		return err
	}

	if err := AssertIsDir(organizedir); err != nil {
		return err
	}

	for _, src := range srcs {
		if err := AssertExists(src); err != nil {
			return err
		}

		ext := strings.TrimLeft(filepath.Ext(src.Filename), ".")

		if ext == "" {
			ext = "Any"
		} else {
			ext = strings.ToUpper(string(ext[0])) + ext[1:]
		}

		destDir := config.Path{
			Filename: filepath.Join(organizedir.Filename, ext),
			Devname:  organizedir.Devname,
			Devices:  organizedir.Devices,
			Active:   organizedir.Active,
			Fsys:     organizedir.Fsys,
		}

		if err := AssertExistsCreateDir(destDir); err != nil {
			return err
		}

		if err := AssertIsDir(destDir); err != nil {
			return err
		}

		if AssertIsDir(src) == nil {
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
