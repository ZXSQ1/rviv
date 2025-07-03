package localfs

import (
	"os"
	"path/filepath"
)

func (local *LocalFs) ListDir(filename string) ([]string, error) {
	filename = filepath.Join(local.prefix, filename)
	entries := []string{}
	rawEntries, err := os.ReadDir(filename)

	if err != nil {
		return nil, err
	}

	for _, rawEntry := range rawEntries {
		entries = append(entries, filename+"/"+rawEntry.Name())
	}

	return entries, nil
}
