package localfs

import "github.com/ZXSQ1/rviv/filesystem"

type LocalFs struct{}

func Init() filesystem.Filesystem {
	return &LocalFs{}
}
