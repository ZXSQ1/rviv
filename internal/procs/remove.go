package procs

import (
	"github.com/ZXSQ1/rviv/internal/config"
	"github.com/ZXSQ1/rviv/internal/info"
)

func Remove(opts config.RemoveOpts) error {
	for _, filename := range opts.Filenames {
		if err := IsExist(filename, nil); err != nil {
			return err
		}

		info.Text(opts.Verbose, "removing file '%s'", ShowPath(filename))

		if !opts.Recursive && IsDir(filename) == nil {
			return info.Error(
				"can not remove directory '%s' (recursive is not specified)",
				ShowPath(filename),
			)
		}

		if IsDir(filename) == nil {
			err := RemoveDir(filename)

			if err != nil {
				return err
			}
		} else {
			err := RemoveFile(filename)

			if err != nil {
				return err
			}
		}

	}

	return nil
}
