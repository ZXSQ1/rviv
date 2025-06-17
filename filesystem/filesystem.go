package filesystem

import (
	"io"
	"io/fs"
)

var (
	// the standard permission for a directory
	DirPerm = fs.FileMode(0755)

	// the standard permission for a regular file
	RegularPerm = fs.FileMode(0644)

	// the standard permission for a symbolic link
	LinkPerm = fs.FileMode(0777)
)

// the filesystem interface that standardizes the operations in all different
// filesystems (e.g. local, FTP, SFTP, WebDav, etc.)
type Filesystem interface {
	// tests for the existence of the file give a path
	IsExist(filename string) bool

	// gives information about the file given a path
	Stat(filename string) (fs.FileInfo, error)

	// creates a file given a path
	Create(filename string) error

	// creates a directory given a path
	CreateDir(filename string) error

	// removes the file given a path
	Remove(filename string) error

	// removes the directory and its contents (if any) given a path
	RemoveDir(filename string) error

	// lists the directory given a path
	ListDir(filename string) ([]string, error)

	// returns an io.ReadWriteCloser for writing and reading a file given a path
	Open(filename string) (io.ReadWriteCloser, error)

	// closes a filesystem (assuming that it is a connection; if it is not a
	// connection, the method returns nil)
	Close() error
}
