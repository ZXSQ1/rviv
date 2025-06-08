package localfs

import "github.com/ZXSQ1/rviv/filesystem"

// an implementation of the Filesystem interface for the local filesystem
type LocalFs struct{}

// returns the LocalFs structure
func Init() filesystem.Filesystem {
	return &LocalFs{}
}
