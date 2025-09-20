package procs

import "github.com/ZXSQ1/rviv/internal/config"

func Sync(opts config.SyncOpts) error {
	if !opts.Src.Active {
		if err := opts.Src.Connect(opts.Verbose); err != nil {
			return err
		}
	}

	if !opts.Dest.Active {
		if err := opts.Dest.Connect(opts.Verbose); err != nil {
			return err
		}
	}

	switch opts.Method {
	case "add":
		return SyncAdd(opts.Src, opts.Dest, opts.Oneway, opts.Verbose)
	case "remove":
		return SyncRemove(opts.Src, opts.Dest, opts.Oneway, opts.Verbose)
	}

	return nil
}
