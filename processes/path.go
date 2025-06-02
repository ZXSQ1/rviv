package processes

import "github.com/ZXSQ1/rviv/filesystem"

type Path struct {
	Filename string
	Filesys  filesystem.Filesystem
}

func NewPath(filename string, filesys filesystem.Filesystem) *Path {
	return &Path{
		Filename: filename,
		Filesys:  filesys,
	}
}
