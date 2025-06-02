package processes

import (
	"os"

	"github.com/ZXSQ1/rviv/filesystem"
)

// continues when paths have permission errors or the like
func ListDir(src *Path, recursive bool) ([]string, error) {
	if !src.Filesys.IsExist(src.Filename) {
		return nil, os.ErrNotExist
	}

	srctype, err := src.Filesys.Type(src.Filename)

	if err != nil {
		return nil, err
	}

	if srctype != filesystem.TypeDir {
		return nil, os.ErrInvalid
	}

	if !recursive {
		return src.Filesys.ListDir(src.Filename)
	}

	entries, err := ListDir(src, false)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		entrytype, err := src.Filesys.Type(entry)

		if err != nil {
			continue
		}

		if entrytype != filesystem.TypeDir {
			entries = append(entries, entry)
		} else {
			newEntries, err := ListDir(
				NewPath(entry, src.Filesys), true,
			)

			if err != nil {
				continue
			}

			entries = append(entries, newEntries...)
		}
	}

	return entries, nil
}
