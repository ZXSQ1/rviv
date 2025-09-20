package procs

import (
	"path/filepath"
	"strings"

	"github.com/ZXSQ1/rviv/config"
)

func OrganizeAlpha(srcs []config.Path, organizedir config.Path,
	verbose bool) error {

	if err := IsDir(organizedir); err != nil {
		return err
	}

	for _, src := range srcs {
		if err := IsExist(src, nil); err != nil {
			return err
		}

		firstLetter := strings.ToUpper(string(
			filepath.Base(src.Filename)[0],
		))

		destDir := CopyPath(
			organizedir, filepath.Join(organizedir.Filename, firstLetter),
		)

		if err := IsExist(destDir, CreateDir); err != nil {
			return err
		}

		if err := IsDir(destDir); err != nil {
			return err
		}

		if IsDir(src) == nil {
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
