package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func Remove(opts config.RemoveOpts) error {
	for _, filename := range opts.Filenames {
		if err := CheckExists(filename); err != nil {
			return err
		}

		info.Text(opts.Verbose, "removing file '%s'", ShowPath(filename))

		if !opts.Recursive && CheckIsDir(filename) == nil {
			return info.Error(
				"can not remove directory '%s' (recursive is not specified)",
				ShowPath(filename),
			)
		}

		if CheckIsDir(filename) == nil {
			err := CheckRemoveDir(filename)

			if err != nil {
				return err
			}
		} else {
			err := CheckRemove(filename)

			if err != nil {
				return err
			}
		}

	}

	return nil
}
