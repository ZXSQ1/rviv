package localfs

import "github.com/ZXSQ1/rviv/filesystem"

type LocalFS struct{}

func Init() filesystem.Filesystem {
	return &LocalFS{}
}
