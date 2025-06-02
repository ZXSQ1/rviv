package filesystem

import "io"

const (
	TypeFile = 0
	TypeDir  = 1
	TypeLink = 2
)

type Filesystem interface {
	IsExist(filename string) (bool, error)
	Type(filename string) (int8, error)
	Size(filename string) (uint64, error)
	CreateFile(filename string) error
	CreateDir(filename string) error
	RemoveFile(filename string) error
	RemoveDir(filename string) error
	ListDir(filename string) ([]string, error)
	GetFile(filename string) (io.ReadWriteCloser, error)
}
