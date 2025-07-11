package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func Copy(opts config.CopyOpts) error {
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
			return info.Error(
				"invalid number of sources to copy for method '%s'",
				opts.Method,
			)
		}

		return CopyFile(
			opts.Srcs[0], opts.Dest, opts.Verbose, func(src, dest string) {
				info.Text(opts.Verbose, "copying file '%s' to '%s'", src, dest)
			},
		)
	case "fd":
		return CopyFiles(opts.Srcs, opts.Dest, opts.Verbose)
	case "dd":
		if len(opts.Srcs) > 1 {
			return info.Error(
				"invalid number of sources to copy for method '%s'",
				opts.Method,
			)
		}

		return CopyDir(opts.Srcs[0], opts.Dest, opts.Verbose)
	case "ad":
		return CopyAll(opts.Srcs, opts.Dest, opts.Verbose)
	}

	return nil
}
