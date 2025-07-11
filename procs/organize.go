package procs

import "github.com/ZXSQ1/rviv/config"

func Organize(opts config.OrganizeOpts) error {
	for _, src := range opts.Srcs {
		if !src.Active {
			if err := src.Connect(opts.Verbose); err != nil {
				return err
			}
		}
	}

	if !opts.OrganizeDir.Active {
		if err := opts.OrganizeDir.Connect(opts.Verbose); err != nil {
			return err
		}
	}

	switch opts.Method {
	case "alpha":
		return OrganizeAlpha(opts.Srcs, opts.OrganizeDir, opts.Verbose)
	case "date":
		return OrganizeDate(
			opts.Srcs, opts.OrganizeDir, opts.Datefmt, opts.Verbose,
		)
	case "ext":
		return OrganizeExt(opts.Srcs, opts.OrganizeDir, opts.Verbose)
	}

	return nil
}
