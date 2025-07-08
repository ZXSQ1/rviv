package localfs

import (
	"os"
	"strings"
)

func (local *LocalFs) ListDir(filename string) ([]string, error) {
	filename = strings.TrimLeft(filename, "/")
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
