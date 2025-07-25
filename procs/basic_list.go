package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/pkg/errors"
)

func List(dirname config.Path) ([]config.Path, error) {
	if err := IsDir(dirname); err != nil {
		return nil, err
	}

	results := []config.Path{}
	rawEntries, err := dirname.Fsys.List(dirname.Filename)

	if err != nil {
		return nil, errors.Wrap(filesystem.ErrListDir, ShowPath(dirname))
	}

	for _, rawEntry := range rawEntries {
		entry := config.Path{
			Devices:  dirname.Devices,
			Devname:  dirname.Devname,
			Filename: rawEntry,
			Fsys:     dirname.Fsys,
			Active:   dirname.Active,
		}

		results = append(results, entry)
	}

	return results, nil
}

func ListRecursive(dirname config.Path) ([]config.Path, error) {
	if err := IsDir(dirname); err != nil {
		return nil, err
	}

	results := []config.Path{}
	entries, err := List(dirname)

	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if err := IsDir(entry); err != nil {
			results = append(results, entry)
			continue
		}

		dirEntries, err := ListRecursive(dirname)

		if err != nil {
			return nil, err
		}

		if len(dirEntries) == 0 {
			results = append(results, entry)
			continue
		}

		results = append(results, dirEntries...)
	}

	return results, nil
}
