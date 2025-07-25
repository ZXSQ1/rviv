package procs

import (
	"io"

	"github.com/ZXSQ1/rviv/config"
	"github.com/ZXSQ1/rviv/filesystem"
	"github.com/pkg/errors"
)

func Open(filename config.Path, mode filesystem.OpenMode) (
	io.ReadWriteCloser, error) {

	if err := IsRegular(filename); err != nil {
		return nil, err
	}

	fileObj, err := filename.Fsys.Open(filename.Filename, mode)

	if err != nil {
		baseErr := filesystem.ErrClosed

		switch mode {
		case filesystem.ModeWrite:
			baseErr = filesystem.ErrOpenWriting
		case filesystem.ModeRead:
			baseErr = filesystem.ErrOpenReading
		}

		return nil, errors.Wrap(baseErr, ShowPath(filename))
	}

	return fileObj, nil
}
