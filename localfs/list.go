package localfs

import (
	"os"

	"github.com/ZXSQ1/rviv/logging"
)

func (local *LocalFs) ListDir(filename string) ([]string, error) {
	entries := []string{}
	rawEntries, err := os.ReadDir(filename)
	logging.ReportErr(err)

	if err != nil {
		return nil, err
	}

	for _, rawEntry := range rawEntries {
		entries = append(entries, filename+"/"+rawEntry.Name())
	}

	return entries, nil
}
