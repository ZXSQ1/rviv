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

	if err := IsExist(opts.Filename, CreateDir); err != nil {
		return nil, err
	}

	if err := IsDir(opts.Filename); err != nil {
		return nil, err
	}

	info.Text(opts.Verbose, "listing directory '%s'", ShowPath(opts.Filename))

	var entries []config.Path
	var err error

	if opts.Recursive {
		entries, err = ListRecursive(opts.Filename)
	} else {
		entries, err = List(opts.Filename)
	}

	if err != nil {
		return nil, info.Error(
			"unable to list directory '%s'", opts.Filename.Filename,
		)
	}

	return entries, nil
}
