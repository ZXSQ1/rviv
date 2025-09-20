package procs

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/pkg/errors"
)

type File struct {
	obj      io.ReadWriteCloser
	mode     filesystem.OpenMode
	filename config.Path
}

func (file *File) Write(p []byte) (n int, err error) {
	if file.mode != filesystem.ModeWrite {
		return 0, errors.Wrap(filesystem.ErrModeWrite, ShowPath(file.filename))
	}

	n, err = file.obj.Write(p)

	if err != nil {
		if n == 0 {
			return n, errors.Wrap(filesystem.ErrEOF, ShowPath(file.filename))
		}

		return n, errors.Wrap(filesystem.ErrWriteFile, ShowPath(file.filename))
	}

	return n, nil
}

func (file *File) Read(p []byte) (n int, err error) {
	if file.mode != filesystem.ModeRead {
		return 0, errors.Wrap(filesystem.ErrModeRead, ShowPath(file.filename))
	}

	n, err = file.obj.Read(p)

	if err != nil {
		if n == 0 {
			return n, errors.Wrap(filesystem.ErrEOF, ShowPath(file.filename))
		}

		return n, errors.Wrap(filesystem.ErrReadFile, ShowPath(file.filename))
	}

	return n, nil
}

func (file *File) Close() error {
	if err := file.obj.Close(); err != nil {
		return errors.Wrap(filesystem.ErrCloseFile, ShowPath(file.filename))
	}

	return nil
}
