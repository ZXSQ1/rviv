package filesystem

import (
	"fmt"
)

var (
	ErrNotExist       = fmt.Errorf("no such file or directory")
	ErrExist          = fmt.Errorf("file already exists")
	ErrFileRegular    = fmt.Errorf("file is a regular file")
	ErrFileNotRegular = fmt.Errorf("file not a regular file")
	ErrFileDir        = fmt.Errorf("file is a directory")
	ErrFileNotDir     = fmt.Errorf("file not a directory")
	ErrFileSym        = fmt.Errorf("file is a symbolic link")
	ErrFileNotSym     = fmt.Errorf("file not a symbolic link")
	ErrModeRead       = fmt.Errorf("file opened in read mode only")
	ErrModeWrite      = fmt.Errorf("file opened in write mode only")
	ErrModeClosed     = fmt.Errorf("file closed")
	ErrModeInvalid    = fmt.Errorf("invalid file open mode")
)
