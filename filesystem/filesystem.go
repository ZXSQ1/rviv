package filesystem

import (
	"io"
	"io/fs"
)

const (
	DirPerm     = 0755
	RegularPerm = 0644
	LinkPerm    = 0777
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
}
