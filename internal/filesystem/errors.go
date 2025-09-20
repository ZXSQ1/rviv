package filesystem

import (
	"fmt"
	"io"
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
	ErrClosed         = fmt.Errorf("file closed")
	ErrModeInvalid    = fmt.Errorf("invalid file open mode")

	ErrWriteFile   = fmt.Errorf("unable to write file")
	ErrReadFile    = fmt.Errorf("unable to read file")
	ErrCloseFile   = fmt.Errorf("unable to close file")
	ErrStat        = fmt.Errorf("unable to stat file")
	ErrOpenWriting = fmt.Errorf("unable to open for writing")
	ErrOpenReading = fmt.Errorf("unable to open for reading")
	ErrListDir     = fmt.Errorf("unable to list directory")
	ErrCreateFile  = fmt.Errorf("unable to create regular file")
	ErrCreateDir   = fmt.Errorf("unable to create directory")
	ErrRemoveFile  = fmt.Errorf("unable to remove regular file")
	ErrRemoveDir   = fmt.Errorf("unable to remove directory")
	ErrEOF         = io.EOF

	ErrCopy = fmt.Errorf("unable to copy source to destination")
	ErrMove = fmt.Errorf("unable to move source to destination")
)
