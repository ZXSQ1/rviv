package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func ListDir(opts config.ListOpts) ([]config.Path, error) {
	if !opts.Filename.Active {
		if err := opts.Filename.Connect(opts.Verbose); err != nil {
			return nil, err
		}
	}

	if !opts.Filename.Fsys.IsExist(opts.Filename.Filename) {
		if opts.Filename.Fsys.CreateDir(opts.Filename.Filename) != nil {
			return nil, info.Error(
				"unable to create directory '%s'", ShowPath(opts.Filename),
			)
		}
	} else {
		stat, err := opts.Filename.Fsys.Stat(opts.Filename.Filename)

		if err != nil {
			return nil, info.Error("unable to stat '%s'", ShowPath(opts.Filename))
		}

		if !stat.IsDir() {
			return nil, info.Error(
				"destination file '%s' is not a directory", ShowPath(opts.Filename),
			)
		}
	}

	info.Text(opts.Verbose, "listing directory '%s'", ShowPath(opts.Filename))

	resultEntries := []config.Path{}
	rawEntries, err := opts.Filename.Fsys.ListDir(opts.Filename.Filename)
	entries := []config.Path{}

	if err != nil {
		return nil, info.Error(
			"unable to list directory '%s'", opts.Filename.Filename,
		)
	}

	for _, rawEntry := range rawEntries {
		entries = append(entries, config.Path{
			Filename: rawEntry,
			Devices:  opts.Filename.Devices,
			Devname:  opts.Filename.Devname,
			Active:   opts.Filename.Active,
			Fsys:     opts.Filename.Fsys,
		})
	}

	for _, entry := range entries {
		stat, err := entry.Fsys.Stat(entry.Filename)

		if err != nil {
			return nil, info.Error("unable to stat file '%s'", entry.Filename)
		}

		if (stat.IsDir() && !opts.Recursive) || stat.Mode().IsRegular() {
			resultEntries = append(resultEntries, entry)
		} else {
			recursiveEntries, err := ListDir(config.ListOpts{
				Filename:  entry,
				Recursive: true,
			})

			if err != nil {
				return nil, info.Error(
					"unable to recursively list directory '%s'", entry.Filename,
				)
			}

			if len(recursiveEntries) == 0 {
				resultEntries = append(resultEntries, entry)
			} else {
				resultEntries = append(resultEntries, recursiveEntries...)
			}
		}
	}

	return resultEntries, nil
}
