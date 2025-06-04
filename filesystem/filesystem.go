package filesystem

import (
	"io"
	"io/fs"
)

const (
	TypeDir     = 0
	TypeRegular = 1
	TypeLink    = 2
	TypeNone    = 3
)

var (
	DirPerm     = fs.FileMode(0755)
	RegularPerm = fs.FileMode(0644)
	LinkPerm    = fs.FileMode(0777)
)

type Filesystem interface {
	IsExist(filename string) bool
	Stat(filename string) (fs.FileInfo, error)
	Create(filename string) error
	CreateDir(filename string) error
	Remove(filename string) error
	RemoveDir(filename string) error
	ListDir(filename string) ([]string, error)
	Open(filename string) (io.ReadWriteCloser, error)
	Close() error
}
