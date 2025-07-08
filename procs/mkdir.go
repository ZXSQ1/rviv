package procs

import (
	"path/filepath"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/info"
)

func Mkdir(opts config.MkdirOpts) error {
	for _, filename := range opts.Filenames {
		if !filename.Active {
			if err := filename.Connect(opts.Verbose); err != nil {
				return err
			}
		}

		if filename.Fsys.IsExist(filename.Filename) {
			stat, err := filename.Fsys.Stat(filename.Filename)

			if err != nil {
				return info.Error("unable to stat file '%s'", ShowPath(filename))
			}

			if !stat.IsDir() {
				return info.Error("file '%s' not a directory", ShowPath(filename))
			}
		}

		if opts.Parent {
			parentDirname := config.Path{
				Filename: filepath.Dir(filename.Filename),
				Devname:  filename.Devname,
				Devices:  filename.Devices,
				Active:   filename.Active,
				Fsys:     filename.Fsys,
			}

			if !parentDirname.Fsys.IsExist(parentDirname.Filename) {
				Mkdir(config.MkdirOpts{
					Filenames: []config.Path{parentDirname},
					Parent:    opts.Parent,
					Verbose:   opts.Verbose,
				})
			}
		}

		info.Text(opts.Verbose, "creating directory '%s'", ShowPath(filename))

		if filename.Fsys.CreateDir(filename.Filename) != nil {
			return info.Error(
				"unable to create directory '%s'", ShowPath(filename),
			)
		}
	}

	return nil
}
