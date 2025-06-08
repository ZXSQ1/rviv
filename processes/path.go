package processes

import "github.com/ZXSQ1/rviv/filesystem"

// the filesystem-independent implementation of the path
type Path struct {
	// the path of the file
	Filename string

	// the filesystem from which the file belongs
	Filesys filesystem.Filesystem
}

// the filesystem-independent implementation of a path slice
type Paths struct {
	// the slice of paths to files
	Filenames []string

	// the filesystem from which the file belongs
	Filesys filesystem.Filesystem
}
