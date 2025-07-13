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

	if err := CheckExistsCreateDir(opts.Filename); err != nil {
		return nil, err
	} else if err := CheckIsDir(opts.Filename); err != nil {
		return nil, err
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
		if CheckIsRegular(entry) == nil || (CheckIsDir(entry) == nil &&
			!opts.Recursive) {

			resultEntries = append(resultEntries, entry)
			continue
		}

		recursiveEntries, err := ListDir(config.ListOpts{
			Filename:  entry,
			Recursive: true,
		})

		if err != nil {
			return nil, err
		}

		if len(recursiveEntries) == 0 {
			resultEntries = append(resultEntries, entry)
		} else {
			resultEntries = append(resultEntries, recursiveEntries...)
		}
	}

	return resultEntries, nil
}
