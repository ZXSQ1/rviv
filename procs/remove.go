package procs

import (
	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func Remove(opts config.RemoveOpts) error {
	for _, filename := range opts.Filenames {
		if !filename.Fsys.IsExist(filename.Filename) {
			return info.Error("file '%s' does not exist", ShowPath(filename))
		}

		info.Text(opts.Verbose, "removing file '%s'", ShowPath(filename))
		stat, err := filename.Fsys.Stat(filename.Filename)

		if err != nil {
			return info.Error("unable to stat '%s'", ShowPath(filename))
		}

		if !opts.Recursive && stat.IsDir() {
			return info.Error(
				"can not remove directory '%s' (recursive is not specified)",
				ShowPath(filename),
			)
		}

		if stat.IsDir() {
			err = filename.Fsys.RemoveDir(filename.Filename)

			if err != nil {
				return info.Error("unable to remove directory '%s'", ShowPath(
					filename))
			}
		} else {
			err = filename.Fsys.Remove(filename.Filename)

			if err != nil {
				return info.Error("unable to remove file '%s'", ShowPath(filename))
			}
		}

	}

	return nil
}
