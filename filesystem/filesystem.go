package filesystem

import (
	"io"
	"io/fs"
)

type OpenMode uint8

var (
	PermDir              = fs.FileMode(0755)
	PermRegular          = fs.FileMode(0644)
	PermLink             = fs.FileMode(0777)
	ModeWrite   OpenMode = 0
	ModeRead    OpenMode = 1
	ModeClosed  OpenMode = 2
)

const (
	BufferSize = 1024 * 1000
)

// the filesystem interface that standardizes the operations in all different
// filesystems (e.g. local, FTP, SFTP, WebDav, etc.)
//
// note: some methods depend on other methods; mainly the ones depended on are
// IsExist() and Stat()
type Filesystem interface {
	// tests for the existence of the file give a path
	IsExist(filename string) bool

	// gives information about the file given a path;
	// - Name() gives the base name
	// - Size() gives the size of the regular file and -1 for a directory
	// - Mode() gives the default modes in the filesystem module (e.g. PermDir)
	// - ModTime() gives the modification time
	// - IsDir() checks if the file is a directory
	// - Sys() completely and utterly useless; returns nil always
	Stat(filename string) (fs.FileInfo, error)

	// creates a file given a path; returns an error if the regular file exists
	Create(filename string) error

	// creates a directory given a path; returns an error if the directory exists
	CreateDir(filename string) error

	// removes the file given a path; returns an error if filename refers to a
	// directory
	Remove(filename string) error

	// removes the directory and its contents (if any) given a path; returns an
	// error if filename refers to a regular file
	RemoveDir(filename string) error

	// lists the directory given a path
	ListDir(filename string) ([]string, error)

	// returns an io.ReadWriteCloser for writing and reading a file given a path;
	// the mode is one of WriteMode or ReadMode; methods return -1 on error as n
	// except for EOF (returns 0)
	Open(filename string, mode OpenMode) (io.ReadWriteCloser, error)

	// closes a filesystem (assuming that it is a connection; if it is not a
	// connection, the method returns nil)
	Close() error
}
