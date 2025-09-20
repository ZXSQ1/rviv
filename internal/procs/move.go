package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func Move(opts config.MoveOpts) error {
	for _, src := range opts.Srcs {
		if !src.Active {
			if err := src.Connect(opts.Verbose); err != nil {
				return err
			}
		}
	}

	if !opts.Dest.Active {
		if err := opts.Dest.Connect(opts.Verbose); err != nil {
			return err
		}
	}

	switch opts.Method {
	case "ff":
		if len(opts.Srcs) > 1 {
			info.Error(
				"invalid number of sources to move for method '%s'",
				opts.Method,
			)
		}

		return MoveFile(
			opts.Srcs[0], opts.Dest, opts.Verbose, func(src, dest string) {
				info.Text(opts.Verbose, "moving file '%s' to '%s'", src, dest)
			},
		)
	case "fd":
		return MoveFiles(opts.Srcs, opts.Dest, opts.Verbose)
	case "dd":
		if len(opts.Srcs) > 1 {
			info.Error(
				"invalid number of sources to move for method '%s'",
				opts.Method,
			)
		}

		return MoveDir(opts.Srcs[0], opts.Dest, opts.Verbose)
	case "ad":
		return MoveAll(opts.Srcs, opts.Dest, opts.Verbose)
	}

	return nil
}
