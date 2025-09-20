package localfs

import (
	"path/filepath"
	"slices"
	"strings"

	"github.com/ZXSQ1/rviv/internal/filesystem"
)

func (local *LocalFs) List(filename string) ([]string, error) {
	filename = strings.TrimLeft(filename, "/")
	entries := []string{}
	directory, err := local.fsys.Open(filename)

	if err != nil {
		return nil, err
	}

	dirStat, err := directory.Stat()

	if err != nil {
		return nil, err
	}

	if !dirStat.IsDir() {
		return nil, filesystem.ErrFileNotDir
	}

	rawEntries, err := directory.Readdir(-1)

	if err != nil {
		return nil, err
	}

	for _, rawEntry := range rawEntries {
		entries = append(
			entries, filepath.Join(filename, rawEntry.Name()),
		)
	}

	slices.Sort(entries)

	return entries, nil
}
