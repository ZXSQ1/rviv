package processes

import "github.com/ZXSQ1/rviv/filesystem"

type Path struct {
	Filename string
	Filesys  filesystem.Filesystem
}
